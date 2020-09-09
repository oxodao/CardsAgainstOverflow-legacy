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
	Displays         []*Display `json:"-"`
	ZenMode          bool
	Turn             int
	MaxTurn          int
	Started          bool
	CurrentBlackCard *Card
	DefaultRerollTimeout int

	TurnState        TurnState
	DefaultCountdown int
	CurrentCountdown int `json:"-"`

	SelectedCards []*Card `json:"-"`

	RemainingCards      []*Card `json:"-"`
	RemainingBlackCards []*Card `json:"-"`
	UsedCards           []*Card `json:"-"`
	UsedBlackCards      []*Card `json:"-"`

	JudgeSelection int
	Answers        []*Proposal
	WinningAnswer  *Proposal
	Winner         string

	AvailableDecks []*Deck
}

func (r *Room) HasEnoughCards() bool {
	amtCards := 0
	for _, c := range r.AvailableDecks {
		if c.IsSelected {
			amtCards = amtCards + c.AmtWhite
		}
	}
	return ((len(r.Participants)+1) * 7) <= amtCards
}

// IsReady returns whether the game can start
func (r *Room) IsReady() bool {
	return len(r.Participants) >= 3 && r.HasEnoughCards()
}

// PickCard picks a new white card and put it in the used backlog
func (r *Room) PickCard() *Card {
	amtRemaining := len(r.RemainingCards)

	// If there are no more cards, we refill
	if amtRemaining == 0 {
		// We put back all cards
		for _, c := range r.UsedCards {
			r.RemainingCards = append(r.RemainingCards, c)
		}

		r.UsedCards = []*Card{}

		// And we set back the remaining cards amount
		amtRemaining = len(r.RemainingCards)
	}

	cardPicked := rand.Intn(amtRemaining)
	card := r.RemainingCards[cardPicked]

	r.RemainingCards[cardPicked] = r.RemainingCards[len(r.RemainingCards)-1]
	r.RemainingCards = r.RemainingCards[:len(r.RemainingCards)-1]

	return card
}

// PickBlackCard picks a new black card and put it in the used backlog
func (r *Room) PickBlackCard() *Card {
	amtRemaining := len(r.RemainingBlackCards)

	// If there are no more cards, we refill
	if amtRemaining == 0 {
		for _, v := range r.UsedBlackCards {
			// We don't take the current card of course
			if v.ID != r.CurrentBlackCard.ID {
				r.RemainingBlackCards = append(r.RemainingBlackCards, v)
			}
		}

		r.UsedCards = []*Card{}
		amtRemaining = len(r.RemainingBlackCards)
	}

	cardPicked := rand.Intn(amtRemaining)
	card := r.RemainingBlackCards[cardPicked]

	r.RemainingBlackCards[cardPicked] = r.RemainingBlackCards[len(r.RemainingBlackCards)-1]
	r.RemainingBlackCards = r.RemainingBlackCards[:len(r.RemainingBlackCards)-1]

	r.UsedBlackCards = append(r.UsedBlackCards, card)

	return card
}


// Utils
func GetAllPlayersCardsID(r *Room) []int64 {
	cards := []int64{}
	for _, p := range r.Participants {
		for _, c := range p.Hand {
			if c != nil {
				cards = append(cards, c.ID)
			}
		}
	}
	return cards
}

func Contains(cards []int64, card int64) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}