package model

type Deck struct {
	ID                int64   `db:"id"`
	Title             string  `json:"Title" db:"name"`
	Cards             []*Card `json:"Cards" db:"-"`
	BlackCards        []*Card `json:"BlackCards" db:"-"`
	IsSelected        bool    `json:"IsSelected" db:"-"`
	AmtBlack          int     `db:"amt_black"`
	AmtWhite          int     `db:"amt_white"`
	SelectedByDefault bool    `db:"selected_by_default"`
}
