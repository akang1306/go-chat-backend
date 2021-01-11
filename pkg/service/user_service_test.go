package service_test

import (
	"context"
	"testing"

	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/jwt"
	"github.com/challenge/pkg/service"
)

func TestCanCreateUser(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	expectedID := 0
	username := "user1"
	password := "generic_password"

	user, err := service.CreateUser(ctx, username, password)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if user.ID != expectedID {
		t.Errorf("Error: got id %d but expected %d", user.ID, expectedID)
	}
	if user.Username != username {
		t.Errorf("Error: got username %s but expected %s", user.Username, username)
	}
}

func TestCanNotCreateUserWithUsedUsername(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(ctx, username, password)
	_, err := service.CreateUser(ctx, username, password)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotCreateUserWithEmptyUsername(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username := ""
	password := "generic_password"

	_, err := service.CreateUser(ctx, username, password)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotCreateUserWithEmptyPassword(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username := "user"
	password := ""

	_, err := service.CreateUser(ctx, username, password)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}
