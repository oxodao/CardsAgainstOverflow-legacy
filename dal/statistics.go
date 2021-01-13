package dal

import (
	"fmt"
)

func StartedGame(isZenMode bool) (int64, error) {
	DB := GetDatabase()
	res := DB.QueryRow(`INSERT INTO STAT_GAME (IS_ZEN_MODE) VALUES ($1) RETURNING ID`, isZenMode)
	if res.Err() != nil {
		return -1, res.Err()
	}

	var id int64 = -1
	err := res.Scan(&id)

	return id, err
}

func NextTurn(statID int64) error {
	if statID < 0 {
		return nil
	}

	DB := GetDatabase()
	_, err := DB.Exec(`UPDATE STAT_GAME SET AMT_TURNS = AMT_TURNS + 1 WHERE ID = $1`, statID)
	return err
}

func PickedCard(isBlack bool, statID, cardID int64) error {
	if statID < 0 {
		return nil
	}

	rq := "INSERT INTO %v(GAME, PICKED_CARD) VALUES ($1, $2)"
	if isBlack {
		rq = fmt.Sprintf(rq, "STAT_DRAWN_BLACK_CARD")
	} else {
		rq = fmt.Sprintf(rq, "STAT_DRAWN_CARD")
	}

	DB := GetDatabase()
	_, err := DB.Exec(rq, statID, cardID)

	return err
}