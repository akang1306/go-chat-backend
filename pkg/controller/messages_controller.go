package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// SendMessage send a message from one user to another
func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	request := models.SendMessageRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if r.Context().Value("user") != request.Sender {
		http.Error(w, "Token user does not match sender id.", http.StatusUnauthorized)
		return
	}

	msg, err := h.Service.SendMessage(r.Context(),
		request.Sender, request.Recipient, request.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.RespondJSON(w, msg)
}

// GetMessages get the messages from the logged user to a recipient
func (h Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Retrieve list of Messages
	helpers.RespondJSON(w, []*models.Message{{}})
}
