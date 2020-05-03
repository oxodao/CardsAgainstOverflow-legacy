package model

// GameState represents the full state of a room, to be sent as DTO
type GameState struct {
	User *User
	Room *Room
}

// SetUser prepare the user DTO
func (gs *GameState) SetUser(u *User) {
	var userCopy User = *u
	userCopy.Room = nil
	gs.User = &userCopy
}
