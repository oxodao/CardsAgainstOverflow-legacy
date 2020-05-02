package model

import (
	"math/rand"
)

type Room struct {
	RoomID           string
	Participants     []*User
	ZenMode          bool
	Turn             int
	MaxTurn          int
	Started          bool
	CurrentBlackCard *Card

	SelectedDecks []*Card `json:"-"`
	SelectedCards []*Card `json:"-"`

	RemainingCards      []*Card `json:"-"`
	RemainingBlackCards []*Card `json:"-"`
	UsedCards           []*Card `json:"-"`
	UsedBlackCards      []*Card `json:"-"`

	Answers map[*User][]*Card `json:"-"`
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

	r.UsedCards = append(r.UsedCards, card)

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

	r.UsedBlackCards = append(r.UsedBlackCards, card)

	return card
}
