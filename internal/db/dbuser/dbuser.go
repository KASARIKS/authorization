package dbuser

import (
	"crypto/sha256"
	"encoding/hex"
)

type DbUser struct {
	Nickname string `json:"Nickname"`
	Password string `json:"Password"`
}

func NewDbUser(Nickname string, Password string) *DbUser {
	encryptedPassword := sha256.Sum256([]byte(Password))
	newUser := &DbUser{
		Nickname: Nickname,
		Password: hex.EncodeToString(encryptedPassword[:]),
	}

	return newUser
}
