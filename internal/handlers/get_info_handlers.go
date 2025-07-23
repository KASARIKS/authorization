package handlers

import (
	"errors"
	"io"
	"net/http"
)

func GetUserByNickname(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, errors.New("not supported method").Error(), http.StatusBadRequest)
	}

	nickname := r.URL.Query().Get("nickname")
	gottenUser, err := handlersDb.GetUserByNickname(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(w, gottenUser.Nickname+" "+gottenUser.Password)
}
