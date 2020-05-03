package dto

import "github.com/oxodao/cardsagainstoverflow/model"

// Gamestate returns a gamestate without the user
func GameState(r *model.Room) *model.GameState {
	var roomCopy model.Room = *r
	roomCopy.Participants = []*model.User{}

	for _, i := range r.Participants {
		var participantCopy model.User = *i
		participantCopy.Room = nil

		roomCopy.Participants = append(roomCopy.Participants, &participantCopy)
	}

	return &model.GameState{
		Room: &roomCopy,
	}
}
