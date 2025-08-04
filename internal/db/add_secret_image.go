package db

import (
	"database/sql"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func (db *DB) AddImageByNickname(user *dbuser.DbUser, Image []byte) error {
	err := db.CheckUserPassword(user)

	if err != nil {
		return err
	}

	_, err = db.db.Exec("INSERT INTO SecretImages (Nickname, Image) VALUES (:Nickname, :Image);",
		sql.Named("Nickname", user.Nickname),
		sql.Named("Image", Image))

	return err
}
