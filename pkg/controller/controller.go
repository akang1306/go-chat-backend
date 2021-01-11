package controller

import (
	"context"

	"github.com/challenge/pkg/models"
)

type serviceHandler interface {
	CreateUser(ctx context.Context, username, password string) (*models.User, error)
	Login(ctx context.Context, username, password string) (*models.Login, error)
	Health(ctx context.Context) models.Health
}

// Handler provides the interface to handle different requests
type Handler struct {
	Service serviceHandler
}
