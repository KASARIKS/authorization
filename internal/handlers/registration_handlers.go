package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/json" {
		io.WriteString(w, "There's no registration for json yet, sorry.")
		return
	}

	err := addFromForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func addFromForm(r *http.Request) error {
	encryptedPassword := sha256.Sum256([]byte(r.PostFormValue("Password")))

	newUser := &dbuser.DbUser{
		Nickname: r.PostFormValue("Nickname"),
		Password: hex.EncodeToString(encryptedPassword[:]),
	}

	err := handlersDb.AddUser(*newUser)

	return err
}
