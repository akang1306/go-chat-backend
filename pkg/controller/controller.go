package controller

import (
	"github.com/challenge/pkg/models"
)

type serviceHandler interface {
	CreateUser(username, password string) (models.User, error)
	Login(username, password string) (models.Login, error)
	Health() models.Health
}

// Handler provides the interface to handle different requests
type Handler struct {
	Service serviceHandler
}
