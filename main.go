package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/oxodao/cardsagainstoverflow/model"
)

var users []model.User = make([]model.User, 0)
var rooms []*model.Room = make([]*model.Room, 0)

func main() {
	fmt.Println("CardsAgainstOverflow")

	deck := GetAllBoosters()

	fmt.Println(deck)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			fmt.Println("Can't upgrade connection for this client")
			fmt.Println(err)

			return
		}

		params, _ := url.ParseQuery(r.URL.RawQuery)

		client := &model.User{
			Connection: conn,
			Username:   params["username"][0],
			Room:       nil,
		}

		for i := range users {
			if users[i].Username == client.Username {
				client.SendCommand(model.CommandError, "Un utilisateur est déjà connecté avec ce pseudo")
				client.Connection.Close()
				return
			}
		}

		for i := range rooms {
			if rooms[i].RoomID == params["room"][0] {
				client.Room = rooms[i]
				rooms[i].Participants = append(rooms[i].Participants, client)
				break
			}
		}

		if client.Room == nil {
			newID, _ := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz0123456789", 6)
			client.Room = &model.Room{
				RoomID:       strings.ToUpper(newID),
				Participants: []*model.User{client},
			}
			rooms = append(rooms, client.Room)
			fmt.Printf("Creating a new room named %v\n", client.Room.RoomID)
		}

		users = append(users, *client)

		fmt.Printf("User connected! [Name: %v; Room: %v]\n", client.Username, client.Room.RoomID)
		client.SendCommand(model.CommandConnected, struct {
			Username string
			Color    string
			Room     string
		}{
			client.Username,
			client.Color,
			client.Room.RoomID,
		})
	})

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
