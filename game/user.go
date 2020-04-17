package game

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oxodao/cardsagainstoverflow/model"
)

var Users []model.User = make([]model.User, 0)

func Receive(u *model.User) {
	for {
		_, msg, err := u.Connection.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		cmd := model.Command{}
		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			fmt.Println(err)
			break
		}

		ExecuteCommand(u, cmd)
	}
}

func ExecuteCommand(u *model.User, cmd model.Command) {
	switch cmd.Command {
	case "PING":
		u.LastPing = time.Now()
		break

	case "START_GAME":
		if u.IsAdmin {
			StartGame(u.Room)
		}
		break

	case "SELECT_CARD":
		if !u.IsJudge {
			SelectCard(u, cmd.Arguments)
		}
		break

	case "VOTE":
		if u.IsJudge {
			VoteCard(u, cmd.Arguments)
		}
		break

	default:
		fmt.Printf("Unhandled command from %v: %v\n", u.Username, cmd.Command)
	}
}

func SendCommand(u *model.User, command string, payload interface{}) error {
	content, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = u.Connection.WriteJSON(model.Command{
		Command:   command,
		Arguments: string(content),
	})

	return err
}

func Kick(u *model.User, reason string) {
	fmt.Printf("- Disconnecting %v: %v", u.Username, reason)
	QuitRoom(u, reason)
}

func FillHand(u *model.User) {
	for i := range u.Hand {
		if u.Hand[i] == nil {
			u.Hand[i] = u.Room.PickCard()
		}
	}
}
