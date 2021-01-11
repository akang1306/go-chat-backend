package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	if !IsValidForUser(r, request.Sender) {
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
	recipientValue := r.URL.Query().Get("recipient")
	startValue := r.URL.Query().Get("start")
	limitValue := r.URL.Query().Get("limit")

	recipientID, errRecipient := strconv.ParseInt(recipientValue, 0, 64)
	start, errStart := strconv.Atoi(startValue)
	if errStart != nil || errRecipient != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	limit, errLimit := strconv.Atoi(limitValue)
	if errLimit != nil {
		if limitValue != "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		} else {
			limit = 0
		}
	}

	if !IsValidForUser(r, recipientID) {
		http.Error(w, "Token user does not match recipient id.", http.StatusUnauthorized)
		return
	}

	messages, err := h.Service.GetMessages(r.Context(),
		recipientID, start, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.RespondJSON(w, messages)
}

func IsValidForUser(r *http.Request, userID models.UserID) bool {
	return r.Context().Value("user") == userID
}
