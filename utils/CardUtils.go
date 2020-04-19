package utils

import (
	"github.com/oxodao/cardsagainstoverflow/model"
)

func FilterCard(c [7]*model.Card, f func(*model.Card) bool) [7]*model.Card {
	crd := [7]*model.Card{}
	currCard := 0
	for _, cc := range c {
		if f(cc) {
			crd[currCard] = cc
			currCard = currCard + 1
		}
	}
	return crd
}
