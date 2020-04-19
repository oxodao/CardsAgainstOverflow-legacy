package game

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/oxodao/cardsagainstoverflow/dal"
	"github.com/oxodao/cardsagainstoverflow/dto"
	"github.com/oxodao/cardsagainstoverflow/model"
	"github.com/oxodao/cardsagainstoverflow/utils"
)

var Rooms []*model.Room = make([]*model.Room, 0)

func StartTurn(r *model.Room) {
	fmt.Println("Turn is over. Let's go for another one")
	r.CurrentBlackCard = r.PickBlackCard()

	// If the previous player in the list was judge we set it to the current one
	wasJudge := false
	for _, p := range r.Participants {
		FillHand(p)

		willBeJudge := false
		if wasJudge {
			willBeJudge = true
		}
		wasJudge = p.IsJudge
		p.IsJudge = willBeJudge

		fmt.Printf("Player %v isJudge: %v\n", p.Username, p.IsJudge)
	}

	// If the last player was judge, we set the first player as
	if wasJudge {
		r.Participants[0].IsJudge = true
		fmt.Println("First player is now judge")
	}

	SendPlayerList(r)

	r.SelectedCards = []*model.Card{}
	SendCards(r)

	r.Answers = make(map[*model.User][]*model.Card, len(r.Participants))
}

func StartGame(r *model.Room, decks []int) {
	if r.IsReady() {
		r.Started = true
		r.Participants[0].IsJudge = true

		RoomSelectDecks(r, decks)

		r.CurrentBlackCard = r.PickBlackCard()

		for i := range r.Participants {
			FillHand(r.Participants[i])
		}

		r.SelectedCards = []*model.Card{}

		SendCards(r)
		Broadcast(r, model.CommandStarted, struct{}{})
		SendPlayerList(r)

		r.Answers = make(map[*model.User][]*model.Card, len(r.Participants))
	}
}

func RoomSelectDecks(r *model.Room, selectedDecks []int) error {
	var err error
	var decks []*model.Deck

	if len(selectedDecks) == 0 {
		decks, err = dal.FetchAllDecks() // @TODO: use only what players want
		if err != nil {                  // Use selected decks. Fallback on FatchAllDecks only when no deck are selected
			return err
		}
	} else {

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

func Join(u *model.User, r *model.Room) {
	u.Room = r
	r.Participants = append(r.Participants, u)

	SendPlayerList(r)
	if r.Started {
		FillHand(u)
		SendCardsToUser(u)
	}
}

func QuitRoom(u *model.User, reason string) {
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
			fmt.Println("Removing room #" + room.RoomID)
			Rooms = append(Rooms[:index], Rooms[index+1:]...)
		}
	}

	u.Connection.Close()

	SendPlayerList(room)
}

func SendPlayerList(r *model.Room) {
	Broadcast(r, model.CommandPlayerList, dto.DTOPlayerList(r.Participants))
}

func SendCards(r *model.Room) {
	for i := range r.Participants {
		user := r.Participants[i]
		SendCommand(user, model.CommandUpdateCards, struct {
			Hand      [7]*model.Card
			BlackCard *model.Card
		}{
			Hand:      user.Hand,
			BlackCard: r.CurrentBlackCard,
		})
	}
}

func ReceiveAnswers(u *model.User, argsStr string) {
	args := []model.Card{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		fmt.Println("Bad answers format", err)
		return
	}

	cardsPointers := []*model.Card{}
	for i := range args {
		cardsPointers = append(cardsPointers, &args[i])
	}

	if u.IsJudge {
		r := u.Room
		winner := FindWinnerFromCard(r, cardsPointers[0])
		// If it is not...
		// Oh dear we're in trouble
		if winner != nil {
			winner.Score = winner.Score + 1

			// We remove used cards from player hands

			// For each user we have a set of answers
			for u, a := range r.Answers {
				// For each answer
				for _, currAnswer := range a {
					u.Hand = utils.FilterCard(u.Hand, func(cc *model.Card) bool {
						return currAnswer.ID != cc.ID
					})
				}
			}

			// We clear the answers received
			r.Answers = make(map[*model.User][]*model.Card, 0)

			// Start a new turn
			// No need to broadcast the winner (Client can simply diff from last state)
			// No need to broadcast playerlist because starting a turn sets the judge thus sending the list
			StartTurn(r)
		}
	} else {
		u.Room.Answers[u] = cardsPointers
		for i := range u.Room.Participants {
			judge := u.Room.Participants[i]
			if judge.IsJudge {
				SendCommand(judge, model.CommandSendAnswersList, cardsPointers)
				break
			}
		}
	}
}

func FindWinnerFromCard(r *model.Room, c *model.Card) *model.User {
	for k := range r.Answers {
		for j := range r.Answers[k] {
			if r.Answers[k][j].ID == c.ID {
				return k
			}
		}
	}

	return nil
}

func SendCardsToUser(user *model.User) {
	SendCommand(user, model.CommandUpdateCards, struct {
		Hand      [7]*model.Card
		BlackCard *model.Card
	}{
		Hand:      user.Hand,
		BlackCard: user.Room.CurrentBlackCard,
	})
}

func Broadcast(r *model.Room, cmdTxt string, arguments interface{}) {
	for i := range r.Participants {
		err := SendCommand(r.Participants[i], cmdTxt, arguments)
		if err != nil {
			fmt.Println("Failed to send command: ", err)
		}
	}
}

func GetAmountCardRequired(c *model.Card) int {
	return strings.Count(c.Text, "____")
}
