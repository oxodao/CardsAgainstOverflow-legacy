package game

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/oxodao/cardsagainstoverflow/dal"
	"github.com/oxodao/cardsagainstoverflow/dto"
	"github.com/oxodao/cardsagainstoverflow/model"
	"github.com/oxodao/cardsagainstoverflow/utils"
)

// Rooms is the list of all rooms on the server
var Rooms []*model.Room = make([]*model.Room, 0)

const DefaultCountdown int = 20

// StartTurn starts a turn
func StartTurn(r *model.Room, gameStarting bool) {
	r.CurrentBlackCard = r.PickBlackCard()

	if gameStarting {
		Log(r, "Starting game!")
		for i := range r.Participants {
			FillHand(r.Participants[i])
		}
	} else {
		r.Turn = r.Turn + 1

		// If the game has ended, no need to do another turn
		if r.Turn > r.MaxTurn && !r.ZenMode {
			r.Started = false
			// Sending the last gamestate before concluding the game
			SendGamestateAll(r)

			Log(r, "Game over!")

			return
		}

		if r.ZenMode {
			Log(r, fmt.Sprintf("Turn %v", r.Turn))
		} else {
			Log(r, fmt.Sprintf("Turn %v / %v", r.Turn, r.MaxTurn))
		}

		// If the previous player in the list was judge we set it to the current one
		wasJudge := false
		for _, p := range r.Participants {
			// We drop played cards
			for _, c := range GetPlayedCards(p) {
				for i, currCard := range p.Hand {
					if c != nil && currCard != nil && c.ID == currCard.ID {
						p.Hand[i] = nil
					}
				}
			}

			// Then we pick new ones
			FillHand(p)

			willBeJudge := false
			if wasJudge {
				willBeJudge = true
			}
			wasJudge = p.IsJudge
			p.IsJudge = willBeJudge
			p.SelectedCards = []int{}
		}

		// If the last player was judge, we set the first player as
		if wasJudge {
			r.Participants[0].IsJudge = true
		}
	}

	r.SelectedCards = []*model.Card{}
	r.CurrentCountdown = r.DefaultCountdown

	r.Answers = []*model.Proposal{}
	r.Winner = ""
	r.WinningAnswer = nil

	r.TurnState = model.TurnStatePlayer

	SendGamestateAll(r)
}

// StartGame starts the game
func StartGame(r *model.Room) {
	if r.IsReady() {
		r.Started = true
		r.Turn = 1
		//r.Turn = r.MaxTurn // Debug only

		for i, p := range r.Participants {
			p.IsJudge = i == 0
			p.Score = 0
		}

		if RoomSelectDecks(r) != nil {
			Log(r, "Can't fetch decks!")
		}

		StartTurn(r, true)
		RunCountdown(r, CountdownProcess)
	}
}

func RoomSelectDecks(r *model.Room) error {
	var err error
	var decks []*model.Deck

	selectedDecks := []int64{}
	for _, d := range r.AvailableDecks {
		if d.IsSelected {
			selectedDecks = append(selectedDecks, d.ID)
		}
	}

	if len(selectedDecks) == 0 {
		decks, err = dal.FetchAllDecks()
	} else {
		decks, err = dal.FetchSelectedDecks(selectedDecks)
	}

	if err != nil {
		return err
	}

	err = dal.FetchCardsForDecks(decks)
	if err != nil {
		return err
	}

	r.RemainingCards = []*model.Card{}
	r.RemainingBlackCards = []*model.Card{}
	for i := range decks {
		r.RemainingCards = append(r.RemainingCards, decks[i].Cards...)
		r.RemainingBlackCards = append(r.RemainingBlackCards, decks[i].BlackCards...)
	}

	// We then calculate each amount of answers per card we want
	// Maybe store this into database later but it will do for now
	for i := range r.RemainingBlackCards {
		r.RemainingBlackCards[i].AmtCardRequired = GetAmountCardRequired(r.RemainingBlackCards[i])
	}

	return nil
}

// Join is triggered when someone joins the room
func Join(u *model.User, r *model.Room) {
	u.Room = r
	r.Participants = append(r.Participants, u)

	if r.Started {
		FillHand(u)
	}

	SendGamestateAll(r)

	Log(r, u.Username+" has joined the room.")
}

func QuitRoom(u *model.User, reason string) {
	Log(u.Room, u.Username+" has left the room.")

	// Setting the next player as the current player
	// If the user was admin, setting the next player as admin
	index := -1
	found := false
	wasSet := false
	for i := range u.Room.Participants {
		// If we get into this, this mean:
		//	- The previous participant was the one quitting
		//	- This is not the last participant in the list
		if found {
			u.Room.Participants[i].IsJudge = u.Room.Participants[i].IsJudge || u.Room.Participants[index].IsJudge
			u.Room.Participants[i].IsAdmin = u.Room.Participants[i].IsAdmin || u.Room.Participants[index].IsAdmin
			wasSet = true
			break
		}

		if u.Room.Participants[i].Username == u.Username {
			index = i
			found = true
		}
	}

	// If it was found but not set
	// This mean the quitting participant was the last in the list
	// So the next one is the first one
	if found && !wasSet {
		u.Room.Participants[0].IsJudge = u.Room.Participants[0].IsJudge || u.Room.Participants[index].IsJudge
		u.Room.Participants[0].IsAdmin = u.Room.Participants[0].IsAdmin || u.Room.Participants[index].IsAdmin
	}

	// Telling everyone that the user is quitting
	Broadcast(u.Room, model.CommandDisconnected, struct {
		Username string
		Reason   string
	}{
		Username: u.Username,
		Reason:   reason,
	})

	// Actually quitting the room
	room := u.Room
	if index != -1 {
		room.Participants = append(room.Participants[:index], room.Participants[index+1:]...)
		u.Room = nil
	}

	// If there are no one left in the room, we remove it
	if len(room.Participants) == 0 {
		index := -1

		for i := range Rooms {
			if room.RoomID == Rooms[i].RoomID {
				index = i
				break
			}
		}

		if index >= 0 {
			Log(room, "Removing the room, no one is inside.")
			Rooms = append(Rooms[:index], Rooms[index+1:]...)
		}
	}

	index = -1
	for i := range Users {
		if Users[i].Username == u.Username { // Not that great but meh, should compare on pointer but doesn't seems to work
			index = i
			break
		}
	}

	if index >= 0 {
		Users = append(Users[:index], Users[index+1:]...)
	}

	u.Connection.Close()

	SendGamestateAll(room)
}

func SendGamestate(u *model.User) {
	gs := dto.GameState(u.Room)
	gs.SetUser(u)
	SendCommand(u, model.CommandSetGamestate, gs)
}

func SendGamestateAll(r *model.Room) {
	for _, p := range r.Participants {
		SendGamestate(p)
	}
}

// ReceiveAnswers set the answer for the user
func ReceiveAnswers(u *model.User, argsString string) {
	if u.Room.Turn > u.Room.MaxTurn && !u.Room.ZenMode {
		return
	}

	args := []int{}
	err := json.Unmarshal([]byte(argsString), &args)
	if err != nil {
		Log(u.Room, u.Username+" > Can't parse the received cards!")
		return
	}

	if u.IsJudge {
		if u.Room.TurnState == model.TurnStateJudge {
			// The judge has given his verdict!
			winnerCard := args[0]
			if len(u.Room.Answers) >= winnerCard {
				u.Room.WinningAnswer = u.Room.Answers[winnerCard]
			}

			Broadcast(u.Room, model.CommandJudgeSelection, winnerCard)
		}
	} else {
		if u.Room.TurnState == model.TurnStatePlayer {
			// We can process the request
			u.SelectedCards = args

			canDropCounter := true
			for _, p := range u.Room.Participants {
				if p.IsJudge {
					continue
				}

				amtSelect := 0
				for _, c := range p.SelectedCards {
					if c != -1 {
						amtSelect = amtSelect + 1
					}
				}

				if amtSelect < u.Room.CurrentBlackCard.AmtCardRequired {
					canDropCounter = false
				}
			}

			// We set the final countdown tudoduto tudodutoto
			if canDropCounter && u.Room.CurrentCountdown > 10 {
				u.Room.CurrentCountdown = 10
			}
		}
	}
}

// CountdownProcess is a function called when a countdown reach 0
func CountdownProcess(r *model.Room) {
	// We switch to the next game state
	if r.TurnState == model.TurnStatePlayer || r.TurnState == model.TurnStateJudge {
		if r.TurnState == model.TurnStatePlayer {
			r.TurnState = model.TurnStateJudge
			r.CurrentCountdown = r.DefaultCountdown

			for _, p := range r.Participants {
				// We don't take the judge
				if p.IsJudge {
					continue
				}

				if utils.HasPlayerPlayed(r.CurrentBlackCard, p.SelectedCards) {
					cards := []*model.Card{}
					if p.SelectedCards != nil {
						for _, a := range p.SelectedCards {
							cards = append(cards, p.Hand[a])
						}
					}

					r.Answers = append(r.Answers, &model.Proposal{
						Cards: cards,
						User:  p,
					})
				}
			}

			if len(r.Answers) == 0 {
				StartTurn(r, false)
				return
			}

			utils.Shuffle(r.Answers)

			for i, c := range r.Answers {
				c.ID = i
			}
		} else if r.TurnState == model.TurnStateJudge {
			r.TurnState = model.TurnStateShowWinner
			r.CurrentCountdown = 6

			if r.WinningAnswer == nil {
				// The judge has not answered so we'll give it to a random player
				size := len(r.Answers) - 1
				selected := 0

				// If there are multiple players
				if size > 0 {
					selected = rand.Intn(size)
				}

				r.WinningAnswer = r.Answers[selected]
			}

			r.Winner = r.WinningAnswer.User.Username
			r.WinningAnswer.User.Score = r.WinningAnswer.User.Score + 1
		}

		// Send GameState
		SendGamestateAll(r)

	} else if r.TurnState == model.TurnStateShowWinner {
		StartTurn(r, false)
	}

}

type gotSettings struct {
	SelectedDecks []int64
	MaxTurn int
	ZenMode bool
	DefaultCountdown int
}

func SetSettings(u *model.User, argStr string) {
	settings := gotSettings{}
	err := json.Unmarshal([]byte(argStr), &settings)
	if err != nil {
		Log(u.Room, fmt.Sprint("Can't parse settings: %v", err))
		return
	}

	for _, v := range u.Room.AvailableDecks {
		v.IsSelected = model.Contains(settings.SelectedDecks, v.ID)
	}
	u.Room.MaxTurn = settings.MaxTurn
	u.Room.ZenMode = settings.ZenMode
	u.Room.DefaultCountdown = settings.DefaultCountdown

	Broadcast(u.Room, model.CommandGotSettings, settings)

}

// SendPlayerList broadcast the playerlist to everyone in a room
func SendPlayerList(r *model.Room) {
	Broadcast(r, model.CommandPlayerList, dto.Participants(r.Participants))
}

// Broadcast sends a command to all users in a room
func Broadcast(r *model.Room, cmdTxt string, arguments interface{}) {
	for i := range r.Participants {
		err := SendCommand(r.Participants[i], cmdTxt, arguments)
		if err != nil {
			fmt.Println("Failed to send command: ", err)
		}
	}
}

func GetPlayedCards(u *model.User) []*model.Card {
	for _, p := range u.Room.Answers {
		if p.User == u {
			return p.Cards
		}
	}
	return nil
}

// GetAmountCardRequired counts the amount of
func GetAmountCardRequired(c *model.Card) int {
	return strings.Count(c.Text, "____")
}
