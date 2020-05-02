package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/cardsagainstoverflow/model"
)

func FetchAllDecks() ([]*model.Deck, error) {
	DB := GetDatabase()

	var decks []*model.Deck
	rows, err := DB.Queryx("SELECT ID, NAME FROM DECK")
	if err != nil {
		return decks, err
	}

	for rows.Next() {
		deck := &model.Deck{}
		rows.StructScan(&deck)

		decks = append(decks, deck)
	}

	return decks, nil
}

func FetchCardsForDecks(decks []*model.Deck) error {
	DB := GetDatabase()
	for i := range decks {
		cards, err := FetchCardsForDeck(DB, decks[i].ID, false)
		if err != nil {
			return err
		}

		decks[i].Cards = cards

		cards, err = FetchCardsForDeck(DB, decks[i].ID, true)
		if err != nil {
			return err
		}

		decks[i].BlackCards = cards

	}

	return nil
}

func FetchCardsForDeck(DB *sqlx.DB, deck int64, isBlack bool) ([]*model.Card, error) {
	var cards []*model.Card
	i := 0
	if isBlack {
		i = 1
	}

	rows, err := DB.Queryx(`
	SELECT ID, TEXT
	FROM CARD
	WHERE ID IN (
		SELECT CARD_ID FROM CARD_DECK WHERE DECK_ID = ?
	)
	AND IS_BLACK_CARD = ?`, deck, i)
	if err != nil {
		return cards, err
	}

	for rows.Next() {
		card := &model.Card{}
		rows.StructScan(&card)

		cards = append(cards, card)
	}

	return cards, nil
}