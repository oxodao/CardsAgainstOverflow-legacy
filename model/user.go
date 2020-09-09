package model

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// User reprensents a user
type User struct {
	Username      string
	IsAdmin       bool
	Room          *Room `json:"-"`
	Score         int
	Hand          [7]*Card
	RerollTimeout int

	LastWizz  int `json:"-"`
	HasPlayed bool

	IsJudge       bool
	SelectedCards []int `json:"-"`

	MutexWS    *sync.Mutex     `json:"-"`
	Connection *websocket.Conn `json:"-"`
	LastPing   time.Time       `json:"-"`
}

type Display struct {
	Room       *Room           `json:"-"`
	MutexWS    *sync.Mutex     `json:"-"`
	Connection *websocket.Conn `json:"-"`
	LastPing   time.Time       `json:"-"`
}
