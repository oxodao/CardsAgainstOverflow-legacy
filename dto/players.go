package dto

import (
	"github.com/oxodao/cardsagainstoverflow/model"
)

func DTOPlayerList(users []*model.User) []model.User {
	usersCleaned := []model.User{}

	for i := range users {
		usersCleaned = append(usersCleaned, *users[i])
		usersCleaned[i].Room = nil
	}

	return usersCleaned
}
