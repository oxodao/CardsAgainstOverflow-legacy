package game

import (
	"time"

	"github.com/oxodao/cardsagainstoverflow/model"
)

// RunCountdown starts a timer to automatically switch the
func RunCountdown(r *model.Room, callback func(*model.Room)) {
	// We set the initial countdown to 60 per TurnState
	r.CurrentCountdown = DefaultCountdown
	go Countdown(r, callback)
}

// Countdown creates a countdown for the current room
func Countdown(r *model.Room, callback func(*model.Room)) {
	for {
		time.Sleep(1 * time.Second)
		r.CurrentCountdown = r.CurrentCountdown - 1
		//fmt.Printf("Countdown tick: %v\n", r.CurrentCountdown)
		Broadcast(r, model.CommandCountdown, r.CurrentCountdown)

		if r.CurrentCountdown <= 0 {
			callback(r)
		}
	}
}
