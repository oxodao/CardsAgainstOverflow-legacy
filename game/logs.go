package game

import (
	"fmt"

	"github.com/oxodao/cardsagainstoverflow/model"
)

// Log write in the terminal some text with a prefix for the room
func Log(r *model.Room, txt string) {
	fmt.Printf("[%v] %v\n", r.RoomID, txt)
}
