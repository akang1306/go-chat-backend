package service

import (
	"context"

	"github.com/challenge/pkg/models"
)

func (h Handler) Health(_ context.Context) models.Health {
	return models.Health{Health: "ok"}
}
