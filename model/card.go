package model

type Card struct {
	ID          int64  `json:"ID"`
	Text        string `json:"Text"`
	IsBlackCard bool   `json:"BlackCard"`
}
