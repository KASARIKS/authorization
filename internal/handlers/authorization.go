package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/authorization/internal/db/dbuser"
)

const jwtKey = "oepuqtpowejfvgc;lvmjn290331pq;woerwqje[p]"
const jwtExistingTimeMinutes = 2

func AuthorizationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/authorization.html")
}

// Todo - fix jwt type
func Authorization(w http.ResponseWriter, r *http.Request) {
	err := handlersDb.CheckUserPassword(dbuser.NewDbUser(
		r.PostFormValue("Nickname"),
		r.PostFormValue("Password")))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{
		"Nickname": r.PostFormValue("Nickname"),
		"exp":      time.Now().Add(time.Minute * jwtExistingTimeMinutes).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := jwtToken.SignedString([]byte(jwtKey))
	if err != nil {
		io.WriteString(w, fmt.Sprintf("failed to sign jwt: %s\n", err))
	}

	cookie := http.Cookie{
		Name:     "jwtToken",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Minute * jwtExistingTimeMinutes),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
}
