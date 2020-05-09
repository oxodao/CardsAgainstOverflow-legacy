package main

import (
	"github.com/oxodao/cardsagainstoverflow/dal"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/oxodao/cardsagainstoverflow/dto"
	"github.com/oxodao/cardsagainstoverflow/game"
	"github.com/oxodao/cardsagainstoverflow/model"
)

func ConnectUser(conn *websocket.Conn, params url.Values) {
	client := &model.User{
		Connection: conn,
		Username:   params["username"][0],
		Room:       nil,
		LastPing:   time.Now(),
	}

	for i := range game.Rooms {
		if game.Rooms[i].RoomID == params["room"][0] {
			// If there is already someone in the room with the same username, we don't let him connect
			for _, p := range game.Rooms[i].Participants {
				if p.Username == client.Username {
					game.SendCommand(client, model.CommandCriticalError, "Un utilisateur est déjà connecté avec ce pseudo")
					client.Connection.Close()
					return
				}
			}
			game.Join(client, game.Rooms[i])
			break
		}
	}

	if client.Room == nil && len(params["room"][0]) == 0 {
		r := GenerateNewRoom(client)

		client.IsAdmin = true

		game.Join(client, r)
		game.Rooms = append(game.Rooms, r)
	} else if (client.Room) == nil {
		game.SendCommand(client, model.CommandCriticalError, "Salle inexistante")
		conn.Close()
		return
	}

	gs := dto.GameState(client.Room)
	gs.SetUser(client)

	game.SendCommand(client, model.CommandSetGamestate, gs)

	go game.Receive(client)

	conn.SetCloseHandler(func(code int, text string) error {
		game.QuitRoom(client, "Bye-bye")
		return nil
	})
}

// GenerateNewRoom create a room and set the default parameters
func GenerateNewRoom(client *model.User) *model.Room {
	newID, _ := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz0123456789", 6)
	decks, _ := dal.FetchAllDecks()
	r := &model.Room{
		RoomID:       strings.ToUpper(newID),
		Participants: []*model.User{},
		Started:      false,
		MaxTurn:      10,
		AvailableDecks: decks,
		DefaultRerollTimeout: 4,
		DefaultCountdown: 80,
	}

	game.Rooms = append(game.Rooms, r)
	game.RunCountdown(r, game.CountdownProcess)
	game.Log(r, "Creating the room.")

	return r
}
