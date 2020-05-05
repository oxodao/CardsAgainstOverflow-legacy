package dto

import (
	"github.com/oxodao/cardsagainstoverflow/model"
)

// Participant is a DTO for other users (Not the current player)
func Participant(u *model.User) *model.User {
	var user model.User = *u
	u.Hand = [7]*model.Card{}
	return &user
}

// Participants is a DTO for other users (Not the current player)
func Participants(users []*model.User) []*model.User {
	usersCopied := []*model.User{}

	for _, participant := range users {
		usersCopied = append(usersCopied, Participant(participant))
	}

	return usersCopied
}
