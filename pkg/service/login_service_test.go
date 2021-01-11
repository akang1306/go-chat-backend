package service_test

import (
	"context"
	"testing"

	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/jwt"
	"github.com/challenge/pkg/service"
)

func TestLoginWorksAsExpected(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	expectedID := 0
	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(ctx, username, password)

	login, err := service.Login(ctx, username, password)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if login.UserID != expectedID {
		t.Errorf("Error: got id %d but expected %d", login.UserID, expectedID)
	}
	if login.Token == "" {
		t.Errorf("Error: got empty token")
	}
}

func TestCanNotLoginInvalidUsername(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(ctx, username, password)

	_, err := service.Login(ctx, "user2", password)
	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotLoginInvalidPassword(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(ctx, username, password)

	_, err := service.Login(ctx, username, "invalidpwd")
	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}
