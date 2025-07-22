package db

import (
	"database/sql"

	"github.com/kasariks/authorization/internal/dbUser"
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

func (db *DB) AddUser(dbUser dbUser.DbUser) error {
	_, err := db.db.Exec("INSERT INTO Users (Nickname, Password) VALUES (:Nickname, :Password);",
		sql.Named("Nickname", dbUser.Nickname),
		sql.Named("Password", dbUser.Password))

	return err
}

func (db *DB) GetUserByNickname(Nickname string) (*dbUser.DbUser, error) {
	row := db.db.QueryRow("SELECT Nickname, Password FROM Users WHERE Nickname=:Nickname;",
		sql.Named("Nickname", Nickname))
	gottenUser := &dbUser.DbUser{}
	err := row.Scan(&gottenUser.Nickname, &gottenUser.Password)

	return gottenUser, err
}
