package dal

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/cardsagainstoverflow/model"
)

func FetchAllDecks() ([]*model.Deck, error) {
	DB := GetDatabase()

	var decks []*model.Deck
	rows, err := DB.Queryx("SELECT ID, NAME, AMT_BLACK, AMT_WHITE, SELECTED_BY_DEFAULT FROM DECK")
	if err != nil {
		return decks, err
	}

	for rows.Next() {
		deck := &model.Deck{}
		rows.StructScan(&deck)

		if deck.SelectedByDefault {
			deck.IsSelected = true
		}

		decks = append(decks, deck)
	}

	return decks, nil
}

func FetchSelectedDecks(selected []int64) ([]*model.Deck, error) {
	if len(selected) == 0 {
		return []*model.Deck{}, errors.New("can't use empty deck set")
	}

	DB := GetDatabase()

	var decks []*model.Deck

	rq := "SELECT ID, NAME FROM DECK WHERE ID IN ("
	args := []interface{}{}
	for i, v := range selected {
		rq = rq + "?"
		args = append(args, v)

		if i != len(selected)-1 {
			rq = rq + ", "
		}
	}

	rq = rq + ");"

	rows, err := DB.Queryx(rq, args...)
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
	SELECT c.ID, c.TEXT, cd.DECK_ID as DECK
	FROM CARD c
		INNER JOIN CARD_DECK cd ON c.ID = cd.CARD_ID
	WHERE cd.DECK_ID = ?
		AND c.IS_BLACK_CARD = ?`, deck, i)
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
