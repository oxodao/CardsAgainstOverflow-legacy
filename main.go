package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oxodao/cardsagainstoverflow/dal"
	"github.com/oxodao/cardsagainstoverflow/game"
)

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

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			fmt.Println("Can't upgrade connection for this client")
			fmt.Println(err)

			return
		}

		params, _ := url.ParseQuery(r.URL.RawQuery)

		ConnectUser(conn, params)
	})

	// Disconnecting old players
	go func() {
		maxDelay, _ := time.ParseDuration("-30s")

		for {
			for i := range game.Rooms {
				for p := range game.Rooms[i].Participants {
					client := game.Rooms[i].Participants[p]
					if client.LastPing.Before(time.Now().Add(maxDelay)) {
						game.Kick(client, "Timed out.")
					}
				}
			}

			time.Sleep(30 * time.Second)
		}
	}()

	err = http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
