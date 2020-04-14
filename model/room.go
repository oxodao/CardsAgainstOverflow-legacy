package model

type Room struct {
	RoomID       string
	Participants []*User
	JoinEvent    chan *User
	ZenMode      bool
	Turn         int
}

func (r *Room) Play() {
	//for {

	//}
}

func (r *Room) Join(u *User) {
	u.Room = r
	r.Participants = append(r.Participants, u)

	r.JoinEvent <- u
}
