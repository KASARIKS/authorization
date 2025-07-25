package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func LoadAddSecretImagePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/addimage.html")
}

func AddSecretImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, errors.New("not supported method").Error(), http.StatusBadRequest)
		return
	}

	nickname := r.FormValue("Nickname")
	gottenUser, err := handlersDb.GetUserByNickname(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imgBuf, err := getImageFromForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handlersDb.AddImageByNickname(gottenUser.Nickname, imgBuf.Bytes())
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
