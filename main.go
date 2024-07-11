package main

import (
	cryptorand "crypto/rand"
	"embed"
	"encoding/binary"
	"flag"
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oxodao/cardsagainstoverflow/dal"
	"github.com/oxodao/cardsagainstoverflow/game"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	var b [8]byte
	_, err := cryptorand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	dbCreationFile := flag.String("create_db", "", "File that contains initial decks to create the database")

	flag.Parse()

	if len(*dbCreationFile) > 0 {
		err := dal.InitializeDB(*dbCreationFile)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("CardsAgainstOverflow")

	// Connect to the DB / Init singleton
	dal.GetDatabase()

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			return
		}

		params, _ := url.ParseQuery(r.URL.RawQuery)
		ConnectUser(conn, params)
	})

	http.HandleFunc("/deporte", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			return
		}

		params, _ := url.ParseQuery(r.URL.RawQuery)
		ConnectDeporte(conn, params)
	})

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	subFs, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.FS(subFs)))

	// Disconnecting old players
	go func() {
		maxDelay, _ := time.ParseDuration("-30s")

		for {
			for _, r := range game.Rooms {
				if r != nil {
					for _, p := range r.Participants {
						if p != nil && p.LastPing.Before(time.Now().Add(maxDelay)) {
							game.Kick(p, "Timed out.")
						}
					}
				}
			}

			time.Sleep(30 * time.Second)
		}
	}()

	fmt.Println("-- Ready --")
	err = http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getAllFilenames(fs *embed.FS, path string) (out []string, err error) {
	if len(path) == 0 {
		path = "."
	}
	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}
