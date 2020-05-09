package model

type Card struct {
	ID              int64  `db:"ID"`
	Text            string `db:"TEXT"`
	IsBlackCard     bool   `db:"IS_BLACK_CARD"`
	Deck            int    `db:"DECK"`
	AmtCardRequired int
}
