package service_test

import (
	"testing"

	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/service"
)

func TestCanCreateUser(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db)

	expectedID := 0
	username := "user1"
	password := "generic_password"

	user, err := service.CreateUser(username, password)

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
	service := service.NewService(db)

	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(username, password)
	_, err := service.CreateUser(username, password)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotCreateUserWithEmptyUsername(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db)

	username := ""
	password := "generic_password"

	_, err := service.CreateUser(username, password)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotCreateUserWithEmptyPassword(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db)

	username := "user"
	password := ""

	_, err := service.CreateUser(username, password)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}
