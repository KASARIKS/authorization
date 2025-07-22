package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"

	"github.com/kasariks/authorization/internal/db"
	"github.com/kasariks/authorization/internal/dbUser"
)

var handlersDb *db.DB

func InitHandlers(db *db.DB) {
	handlersDb = db
}

// Display sign up page
func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/registration.html")
}

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

	newUser := &dbUser.DbUser{
		Nickname: r.PostFormValue("Nickname"),
		Password: hex.EncodeToString(encryptedPassword[:]), // Password isn't encrypted
	}

	err := handlersDb.AddUser(*newUser)

	return err
}

func GetUserByNickname(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, errors.New("not supported method").Error(), http.StatusBadRequest)
	}

	nickname := r.URL.Query().Get("Nickname")
	gottenUser, err := handlersDb.GetUserByNickname(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(w, gottenUser.Nickname+" "+gottenUser.Password)
}
