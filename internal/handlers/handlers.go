package handlers

import (
	"io"
	"net/http"
)

// Display sign up page
func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/registration.html")
}

func Registration(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "There's no registration yet, sorry.")
}
