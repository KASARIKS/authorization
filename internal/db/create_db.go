package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite", "db.db")
	if err != nil {
		return nil, err
	}

	internalDatabase := &DB{
		db: db,
	}

	_, tableCheck := db.Query("SELECT Nickname FROM Users LIMIT 1;")
	if tableCheck != nil {
		_, err = internalDatabase.db.Exec("CREATE TABLE Users (" +
			"Nickname TEXT NOT NULL PRIMARY KEY, " +
			"Password TEXT " +
			");")
		if err != nil {
			return nil, err
		}
	}

	_, tableCheck = db.Query("SELECT Nickname FROM SecretImages LIMIT 1;")
	if tableCheck != nil {
		_, err = internalDatabase.db.Exec("CREATE TABLE SecretImages (" +
			"Nickname TEXT NOT NULL, " +
			"Image BLOB, " +
			"FOREIGN KEY (Nickname) REFERENCES Users (Nickname)" +
			");")
		if err != nil {
			return nil, err
		}
	}

	return internalDatabase, nil
}
