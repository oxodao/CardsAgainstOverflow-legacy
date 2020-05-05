package model

import (
	"time"

	"github.com/gorilla/websocket"
)

// User reprensents a user
type User struct {
	Username string
	IsAdmin  bool
	Room     *Room `json:"-"`
	Score    int
	Hand     [7]*Card

	IsJudge       bool
	SelectedCards []int `json:"-"`

	Connection *websocket.Conn `json:"-"`
	LastPing   time.Time       `json:"-"`
}
