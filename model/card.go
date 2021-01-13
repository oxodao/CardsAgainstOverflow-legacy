package model

type Card struct {
	ID              int64  `db:"id"`
	Text            string `db:"text"`
	IsBlackCard     bool   `db:"is_black_card"`
	Deck            int    `db:"deck"`
	AmtCardRequired int
}
