package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oxodao/cardsagainstoverflow/dal"
	"github.com/oxodao/cardsagainstoverflow/game"
)

//go:embed frontend/dist
var frontend embed.FS

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("CardsAgainstOverflow")

	// Connect to the DB / Init singleton
	dal.GetDatabase()

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		params, _ := url.ParseQuery(r.URL.RawQuery)
		ConnectUser(conn, params)
	})

	http.HandleFunc("/deporte", func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsUpgrader.Upgrade(w, r, nil)
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
