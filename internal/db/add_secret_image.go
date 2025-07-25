package db

import (
	"database/sql"
	"errors"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func (db *DB) AddImageByNickname(user *dbuser.DbUser, Image []byte) error {
	realUser, err := db.GetUserByNickname(user.Nickname)
	if err != nil {
		return err
	}

	if realUser.Password != user.Password {
		return errors.New("wrong password")
	}

	_, err = db.db.Exec("INSERT INTO SecretImages (Nickname, Image) VALUES (:Nickname, :Image);",
		sql.Named("Nickname", realUser.Nickname),
		sql.Named("Image", Image))

	return err
}
