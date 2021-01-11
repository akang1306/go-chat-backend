package controller

import (
	"context"

	"github.com/challenge/pkg/models"
)

type serviceHandler interface {
	CreateUser(ctx context.Context, username, password string) (models.UserID, error)
	Login(ctx context.Context, username, password string) (*models.Login, error)
	Health(ctx context.Context) models.Health
	SendMessage(ctx context.Context, sender, recipient models.UserID, content models.MessageContent) (*models.MessageInfo, error)
	GetMessages(ctx context.Context, recipient models.UserID, start, limit int) ([]*models.Message, error)
}

// Handler provides the interface to handle different requests
type Handler struct {
	Service serviceHandler
}
