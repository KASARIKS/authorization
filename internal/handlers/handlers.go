package handlers

import (
	"net/http"

	"github.com/kasariks/authorization/internal/db"
)

var handlersDb *db.DB

func InitHandlers(db *db.DB) {
	handlersDb = db
}

// Display sign up page
func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/registration.html")
}
