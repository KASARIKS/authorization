package db

import (
	"database/sql"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func (db *DB) AddUser(dbUser dbuser.DbUser) error {
	_, err := db.db.Exec("INSERT INTO Users (Nickname, Password) VALUES (:Nickname, :Password);",
		sql.Named("Nickname", dbUser.Nickname),
		sql.Named("Password", dbUser.Password))

	return err
}
