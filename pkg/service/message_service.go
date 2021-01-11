package service

import (
	"context"

	"github.com/challenge/pkg/models"
)

// SendMessage attempts to store msg to db and returns created message model, else returns error.
func (h Handler) SendMessage(ctx context.Context, sender, recipient int, messageContent models.MessageContent) (*models.Message, error) {
	return h.db.AddMessage(ctx, sender, recipient, messageContent)
}
