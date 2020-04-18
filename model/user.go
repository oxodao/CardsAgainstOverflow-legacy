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

	IsJudge bool

	Connection *websocket.Conn
	LastPing   time.Time
}
