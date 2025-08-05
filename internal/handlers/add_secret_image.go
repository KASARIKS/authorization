package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/authorization/internal/jwtinfo"
)

func LoadAddSecretImagePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/addimage.html")
}

func AddSecretImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, errors.New("not supported method").Error(), http.StatusBadRequest)
		return
	}

	authCookie := r.Cookies()[0]
	authTokenString := authCookie.Value

	authToken, err := jwtinfo.ParseToken(authTokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !authToken.Valid {
		http.Error(w, errors.New("token isn't valid").Error(), http.StatusBadRequest)
		return
	}

	claims, ok := authToken.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, errors.New("invalid token").Error(), http.StatusBadRequest)
		return
	}

	gottenNickname := claims["Nickname"].(string)
	gottenUser, err := handlersDb.GetUserByNickname(gottenNickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	imgBuf, err := getImageFromForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handlersDb.AddImageByNickname(gottenUser, imgBuf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func getImageFromForm(r *http.Request) (bytes.Buffer, error) {
	file, _, err := r.FormFile("SecretImage")
	if err != nil {
		return bytes.Buffer{}, err
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return buf, nil
}
