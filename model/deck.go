package model

type Deck struct {
	ID         int64   `db:"ID"`
	Title      string  `json:"Title" db:"NAME"`
	Cards      []*Card `json:"Cards"`
	BlackCards []*Card `json:"BlackCards"`
	IsSelected bool    `json:"IsSelected" db:"-"`
	SelectedByDefault bool
}
