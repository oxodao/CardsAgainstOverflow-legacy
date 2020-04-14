package model

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	Color    string
	Room     *Room
	Score    int
	Hand     [7]Card

	Connection *websocket.Conn
}

func (u *User) Receive() {
	for {
		_, msg, err := u.Connection.ReadMessage()
		if err != nil {
			break
		}

		cmd := Command{}
		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			break
		}

		u.ExecuteCommand(cmd)
	}
}

func (u *User) ExecuteCommand(cmd Command) {
	// Nothing to do yet
	fmt.Printf("Received command from %v: %v", u.Username, cmd.Command)
}

func (u *User) SendCommand(command string, payload interface{}) error {
	content, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = u.Connection.WriteJSON(Command{
		Command:   command,
		Arguments: string(content),
	})

	return err
}
