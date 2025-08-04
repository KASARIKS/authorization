package db

import (
	"database/sql"
	"errors"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func (db *DB) GetUserByNickname(Nickname string) (*dbuser.DbUser, error) {
	row := db.db.QueryRow("SELECT Nickname, Password FROM Users WHERE Nickname=:Nickname;",
		sql.Named("Nickname", Nickname))
	gottenUser := &dbuser.DbUser{}
	err := row.Scan(&gottenUser.Nickname, &gottenUser.Password)

	return gottenUser, err
}

// return nil if everything okay
func (db *DB) CheckUserPassword(inputUser *dbuser.DbUser) error {
	realUser, err := db.GetUserByNickname(inputUser.Nickname)
	if err != nil {
		return err
	}

	if realUser.Password != inputUser.Password {
		return errors.New("wrong password")
	}

	return nil
}
