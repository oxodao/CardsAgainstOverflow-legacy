package dto

import (
	"github.com/oxodao/cardsagainstoverflow/model"
)

func DTOPlayerList(users []*model.User) []*model.User {
	usersCleaned := []*model.User{}

	for i := range users {
		usersCleaned = append(usersCleaned, DTOUser(users[i]))
	}

	return usersCleaned
}

func DTOUser(u *model.User) *model.User {
	var user model.User = *u
	user.Room = nil
	return &user
}

type modelConnection struct {
	User *model.User
	Room string
}

func DTOConnection(u *model.User) modelConnection {
	user := DTOUser(u)
	return modelConnection{
		User: user,
		Room: u.Room.RoomID,
	}
}
