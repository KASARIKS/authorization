package db

import (
	"database/sql"
)

// Maybe should be checked input password
func (db *DB) AddImageByNickname(Nickname string, Image []byte) error {
	user, err := db.GetUserByNickname(Nickname)
	if err != nil {
		return err
	}

	_, err = db.db.Exec("INSERT INTO SecretImages (Nickname, Image) VALUES (:Nickname, :Image);",
		sql.Named("Nickname", user.Nickname),
		sql.Named("Image", Image))

	return err
}
