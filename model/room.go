package model

import (
	"math/rand"
)

type Room struct {
	RoomID           string
	Participants     []*User
	JoinEvent        chan *User
	ZenMode          bool
	Turn             int
	Started          bool
	CurrentBlackCard *Card

	SelectedDecks []*Card
	SelectedCards []*Card

	RemainingCards      []*Card
	RemainingBlackCards []*Card
}

func (r *Room) IsReady() bool {
	return len(r.Participants) >= 3
}

func (r *Room) PickCard() *Card {
	i := len(r.RemainingCards)
	// If i == 0 => No more cards available
	// We'll do something later when we implement zen mode
	if i == 0 {
		return nil
	}

	cardPicked := rand.Intn(i)
	card := r.RemainingCards[cardPicked]

	r.RemainingCards[cardPicked] = r.RemainingCards[len(r.RemainingCards)-1]
	r.RemainingCards = r.RemainingCards[:len(r.RemainingCards)-1]
	return card

}

func (r *Room) PickBlackCard() *Card {
	i := len(r.RemainingBlackCards)

	// If i == 0 => No more cards available
	// We'll do something later when we implement zen mode
	if i == 0 {
		return nil
	}

	cardPicked := rand.Intn(i)
	card := r.RemainingBlackCards[cardPicked]

	r.RemainingBlackCards[cardPicked] = r.RemainingBlackCards[len(r.RemainingBlackCards)-1]
	r.RemainingBlackCards = r.RemainingBlackCards[:len(r.RemainingBlackCards)-1]

	return card
}
