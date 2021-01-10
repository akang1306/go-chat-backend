package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/helpers"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login authenticates a user and returns a token
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	login, err := h.Service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.RespondJSON(w, login)
}
