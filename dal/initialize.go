package dal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oxodao/cardsagainstoverflow/model"
)

type jsonInputFile struct {
	Decks []model.Deck
}

var dbSingleton *sqlx.DB

func GetDatabase() *sqlx.DB {
	var err error = nil
	if dbSingleton == nil {
		dbSingleton, err = sqlx.Connect("sqlite3", "CAO.db")
		if err != nil {
			panic(err)
		}
	}

	return dbSingleton
}

func InitializeDB(input string) error {
	dbExists := false
	if _, err := os.Stat("CAO.db"); err == nil {
		fmt.Println("Database already exists!")
		fmt.Println("Data will be added to the current DB")
		dbExists = true
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Println("The input file doesn't exists")
		return nil
	}

	jsonInput := jsonInputFile{}

	ctt, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}

	err = json.Unmarshal(ctt, &jsonInput)
	if err != nil {
		return err
	}

	var decks []model.Deck = jsonInput.Decks

	if !dbExists {
		fmt.Println("Creating database...")
	}
	db, err := sqlx.Connect("sqlite3", "./CAO.db")
	if err != nil {
		return err
	}

	if !dbExists {
		fmt.Println("- Initial script executing...")
		err = createDB(db)
		if err != nil {
			return err
		}
	}

	for i := range decks {
		err = insertDeck(db, &decks[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func insertDeck(DB *sqlx.DB, deck *model.Deck) error {
	fmt.Println("\t- Inserting deck '" + deck.Title + "'")
	_, err := DB.Exec("INSERT INTO DECK (NAME) VALUES (?)", deck.Title)
	if err != nil {
		return err
	}

	deck.ID, err = lastInsertedId(DB)
	if err != nil {
		return err
	}

	for i := range deck.Cards {
		fmt.Println("\t\t- Inserting card '" + deck.Cards[i].Text + "'")
		cardID, err := insertCard(DB, deck.Cards[i])
		if err != nil {
			return err
		}

		_, err = DB.Exec("INSERT INTO CARD_DECK (CARD_ID, DECK_ID) VALUES (?, ?)", cardID, deck.ID)
		if err != nil {
			return err
		}
	}

	return nil

}

func insertCard(DB *sqlx.DB, card *model.Card) (int64, error) {
	_, err := DB.Exec("INSERT INTO CARD(TEXT, IS_BLACK_CARD) VALUES (?, ?)", card.Text, card.IsBlackCard)
	if err != nil {
		return -1, nil
	}

	return lastInsertedId(DB)
}

func lastInsertedId(DB *sqlx.DB) (int64, error) {

	row := DB.QueryRowx("SELECT last_insert_rowid()")

	if row.Err() != nil {
		return -1, row.Err()
	}

	var id int64
	row.Scan(&id)

	return id, nil

}

func createDB(DB *sqlx.DB) error {
	rq := `
	CREATE TABLE DECK (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		NAME VARCHAR
	);

	CREATE TABLE CARD (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		TEXT VARCHAR,
		IS_BLACK_CARD INTEGER
	);

	CREATE TABLE CARD_DECK (
		CARD_ID INTEGER,
		DECK_ID INTEGER,
		PRIMARY KEY (CARD_ID, DECK_ID),
		FOREIGN KEY(CARD_ID) REFERENCES CARD(ID),
		FOREIGN KEY(DECK_ID) REFERENCES DECK(ID)
	);`

	_, err := DB.Exec(rq)
	return err
}
