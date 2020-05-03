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
