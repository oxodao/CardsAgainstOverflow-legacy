package model

import (
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	IsAdmin  bool
	Room     *Room
	Score    int
	Hand     [7]*Card

	IsJudge      bool
	SelectedCard *Card `json:"-"` // We don't want to send what other answers

	Connection *websocket.Conn
	LastPing   time.Time
}
