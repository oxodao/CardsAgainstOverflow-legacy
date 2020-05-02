package dto

import "github.com/oxodao/cardsagainstoverflow/model"

// Gamestate is the state of the game at a certain moment
func Gamestate(u *model.User) *model.User {
	var userCopy model.User = *u

	var roomCopy model.Room = *u.Room
	roomCopy.Participants = []*model.User{}

	userCopy.Room = &roomCopy

	for _, i := range u.Room.Participants {
		var participantCopy model.User = *i
		participantCopy.Room = nil

		roomCopy.Participants = append(roomCopy.Participants, &participantCopy)
	}

	return &userCopy
}
