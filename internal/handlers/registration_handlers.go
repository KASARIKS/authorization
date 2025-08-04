package handlers

import (
	"io"
	"net/http"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func RegistrationPage(w http.ResponseWriter, r *http.Request) {
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
	newUser := dbuser.NewDbUser(r.PostFormValue("Nickname"), r.PostFormValue("Password"))
	err := handlersDb.AddUser(*newUser)

	return err
}
