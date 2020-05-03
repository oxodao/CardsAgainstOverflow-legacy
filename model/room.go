package model

import (
	"math/rand"
)

// TurnState represent the current state in a turn
type TurnState int

// TurnStatePlayer is set when the player can send answers
const TurnStatePlayer TurnState = 0

// TurnStateJudge is set when no one except the judge can vote
const TurnStateJudge TurnState = 1

// TurnStateShowWinner is set during a short period of time when the judge has sent its answer to be displayed to everyone
const TurnStateShowWinner TurnState = 2

// Room represents a room
type Room struct {
	RoomID           string
	Participants     []*User
	ZenMode          bool
	Turn             int
	MaxTurn          int
	Started          bool
	CurrentBlackCard *Card

	TurnState        TurnState
	CurrentCountdown int `json:"-"`

	SelectedDecks []*Card `json:"-"`
	SelectedCards []*Card `json:"-"`

	RemainingCards      []*Card `json:"-"`
	RemainingBlackCards []*Card `json:"-"`
	UsedCards           []*Card `json:"-"`
	UsedBlackCards      []*Card `json:"-"`

	Answers       []*Proposal
	WinningAnswer *Proposal
	Winner        string
}

// IsReady returns whether the game can start
func (r *Room) IsReady() bool {
	return len(r.Participants) >= 3
}

// PickCard picks a new white card and put it in the used backlog
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

// PickBlackCard picks a new black card and put it in the used backlog
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
