package utils

import (
	"math/rand"
	"time"

	"github.com/oxodao/cardsagainstoverflow/model"
)

// Shuffle shuffles an array of users
func Shuffle(vals []*model.Proposal) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}

	for i, p := range vals {
		p.ID = i
	}
}

// FilterNegative removes everything thats less than 0
func FilterNegative(val []int) []int {
	vals := []int{}
	for _, p := range val {
		if p != -1 {
			vals = append(vals, p)
		}
	}
	return vals
}

// HasPlayerPlayed returns whether the player gave a correct proposal
func HasPlayerPlayed(blackCard *model.Card, values []int) bool {
	if len(values) != blackCard.AmtCardRequired {
		return false
	}

	for _, val := range values {
		if val == -1 {
			return false
		}
	}

	return true
}
