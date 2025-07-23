package db

import (
	"database/sql"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func (db *DB) GetUserByNickname(Nickname string) (*dbuser.DbUser, error) {
	row := db.db.QueryRow("SELECT Nickname, Password FROM Users WHERE Nickname=:Nickname;",
		sql.Named("Nickname", Nickname))
	gottenUser := &dbuser.DbUser{}
	err := row.Scan(&gottenUser.Nickname, &gottenUser.Password)

	return gottenUser, err
}
