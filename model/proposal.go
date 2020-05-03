package model

// Proposal is one proposal from a user
type Proposal struct {
	ID    int
	Cards []*Card
	User  *User `json:"-"`
}
