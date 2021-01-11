package service

import (
	"context"

	"github.com/challenge/pkg/models"
)

var defaultLimit int = 100

// SendMessage attempts to store msg to db and returns created message model, else returns error.
func (h Handler) SendMessage(ctx context.Context, sender, recipient models.UserID, messageContent models.MessageContent) (*models.MessageInfo, error) {
	return h.db.AddMessage(ctx, sender, recipient, messageContent)
}

// GetMessages attempts to get all messages from db, else returns error.
func (h Handler) GetMessages(ctx context.Context, recipientID models.UserID, start, limit int) ([]*models.Message, error) {
	// Invalid limit values are assumed as default
	if limit <= 0 {
		limit = defaultLimit
	}
	return h.db.GetMessages(ctx, recipientID, start, limit)
}
