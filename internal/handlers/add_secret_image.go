package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/kasariks/authorization/internal/db/dbuser"
)

func LoadAddSecretImagePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/addimage.html")
}

func AddSecretImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, errors.New("not supported method").Error(), http.StatusBadRequest)
		return
	}

	gottenUser := dbuser.NewDbUser(r.FormValue("Nickname"), r.FormValue("Password"))

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
