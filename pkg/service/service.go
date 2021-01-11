package service

import (
	"context"

	"github.com/challenge/pkg/models"
)

type database interface {
	AddUser(ctx context.Context, username string, password []byte) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	AddMessage(ctx context.Context, sender, recipient int, messageContent models.MessageContent) (*models.Message, error)
	GetMessages(ctx context.Context, sender, start, limit int) ([]*models.Message, error)
}

type tokenManager interface {
	TokenForUser(*models.User) string
}

// Handler provides the interface to handle different controller requests
type Handler struct {
	db           database
	tokenManager tokenManager
}

func NewService(db database, tokenManager tokenManager) Handler {
	return Handler{db: db, tokenManager: tokenManager}
}
