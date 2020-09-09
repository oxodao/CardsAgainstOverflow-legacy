package game

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oxodao/cardsagainstoverflow/model"
)

func Reroll(u *model.User) {
	if u.RerollTimeout == 0 {
		u.Hand = [7]*model.Card{}
		u.SelectedCards = []int{}
		u.RerollTimeout = u.Room.DefaultRerollTimeout+1

		FillHand(u)
		SendGamestate(u)
	}
}

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

func ReceiveDisplay(d *model.Display) {
	for {
		_, msg, err := d.Connection.ReadMessage()
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

		ExecuteDisplayCommand(d, cmd)
	}
}

// ExecuteCommand parses and execute a command sent by a user
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

	case "SEND_SELECTION":
		ReceiveAnswers(u, cmd.Arguments)
		break

	case "SKIP_COUNTDOWN":
		if u.IsJudge {
			u.Room.CurrentCountdown = 0
		}
		break

	case "SET_SETTINGS":
		if u.IsAdmin {
			SetSettings(u, cmd.Arguments)
		}
		break;

	case "REROLL":
		if !u.IsJudge {
			Reroll(u)
		}
		break

	case "WIZZ":
		if u.LastWizz == 0 {
			Broadcast(u.Room, "WIZZ", u.Username)
			u.LastWizz = 15
		}
		break

	default:
		fmt.Printf("Unhandled command from %v: %v\n", u.Username, cmd.Command)
	}
}

// ExecuteDisplayCommand parses and execute a command sent by a user
func ExecuteDisplayCommand(d *model.Display, cmd model.Command) {
	switch cmd.Command {
	case "PING":
		d.LastPing = time.Now()
		break

	}
}

// SendCommand sends a command to a user
func SendCommand(u *model.User, command string, payload interface{}) error {
	content, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	u.MutexWS.Lock()
	err = u.Connection.WriteJSON(model.Command{
		Command:   command,
		Arguments: string(content),
	})
	u.MutexWS.Unlock()

	return err
}

// SendCommand sends a command to a user
func SendDisplayCommand(d *model.Display, command string, payload interface{}) error {
	content, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	d.MutexWS.Lock()
	err = d.Connection.WriteJSON(model.Command{
		Command:   command,
		Arguments: string(content),
	})
	d.MutexWS.Unlock()

	return err
}

// Kick disconnects a user from the server
func Kick(u *model.User, reason string) {
	Log(u.Room, u.Username+" > Disconnecting ("+reason+")")
	QuitRoom(u, reason)
}

// FillHand refill the user's hand after a turn
func FillHand(u *model.User) {
	for i := range u.Hand {
		if u.Hand[i] == nil {
			u.Hand[i] = u.Room.PickCard()
		}
	}
}
