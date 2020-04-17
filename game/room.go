package game

import (
	"encoding/json"
	"fmt"

	"github.com/oxodao/cardsagainstoverflow/dal"
	"github.com/oxodao/cardsagainstoverflow/dto"
	"github.com/oxodao/cardsagainstoverflow/model"
)

var Rooms []*model.Room = make([]*model.Room, 0)

func StartTurn(r *model.Room) {
	r.CurrentBlackCard = r.PickBlackCard()

	// If the previous player in the list was judge we set it to the current one
	wasJudge := false
	for i := range r.Participants {
		FillHand(r.Participants[i])
		if wasJudge {
			r.Participants[i].IsJudge = true
		}
		wasJudge = r.Participants[i].IsJudge
		r.Participants[i].IsJudge = false
	}

	r.SelectedCards = []*model.Card{}

	SendCards(r)

	// If the last player was judge, we set the first player as
	if wasJudge {
		r.Participants[0].IsJudge = true
	}
}

func StartGame(r *model.Room) {
	if r.IsReady() {
		r.Started = true
		r.Participants[0].IsJudge = true

		RoomSelectDecks(r)

		r.CurrentBlackCard = r.PickBlackCard()

		for i := range r.Participants {
			FillHand(r.Participants[i])
		}

		r.SelectedCards = []*model.Card{}

		SendCards(r)
		Broadcast(r, model.CommandStarted, struct{}{})
		SendPlayerList(r)
	}
}

func RoomSelectDecks(r *model.Room) error {
	var err error
	var decks []*model.Deck

	if r.SelectedDecks == nil {
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

type cardSelectionRequest struct {
	SelectedCard int64
}

func SelectCard(u *model.User, card string) {
	fmt.Println("Set selected: ", card)
	var csr cardSelectionRequest
	err := json.Unmarshal([]byte(card), &csr)
	if err != nil {
		return
	}

	foundCard, _ := FindCard(u.Room, csr.SelectedCard)
	u.Room.SelectedCards = append(u.Room.SelectedCards, foundCard)

	if len(u.Room.SelectedCards) == (len(u.Room.Participants) - 1) {
		// Everyone has sent his answers

		// @TODO: Add timer if the judge doesn't vote in 60 seconds
		// Choose a random winner
	}

	// We find the judge and send him all selected cards
	for i := range u.Room.Participants {
		u := u.Room.Participants[i]
		if u.IsJudge {
			// We send the cards
			SendCommand(u, model.CommandSendVoteCard, u.Room.SelectedCards)
			break
		}
	}
}

func FindCard(r *model.Room, card int64) (*model.Card, *model.User) {
	for i := range r.Participants {
		p := r.Participants[i]
		for j := range p.Hand {
			if p.Hand[j].ID == card {
				return p.Hand[j], p
			}
		}
	}

	return nil, nil
}

func VoteCard(u *model.User, card string) {

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
