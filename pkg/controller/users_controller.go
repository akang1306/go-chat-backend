package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/helpers"
)

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := userRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id, err := h.Service.CreateUser(userRequest.Username, userRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.RespondJSON(w, id)
}
