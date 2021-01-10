package service

import (
	"github.com/challenge/pkg/jwt"
	"github.com/challenge/pkg/models"
)

type database interface {
	AddUser(user models.User) (int, error)
	GetUserByUsername(username string) (*models.User, error)
}

type tokenManager interface {
	TokenForUser(*models.User) string
}

// Handler provides the interface to handle different controller requests
type Handler struct {
	db           database
	tokenManager tokenManager
}

func NewService(db database) Handler {
	return Handler{db: db, tokenManager: jwt.New()}
}
