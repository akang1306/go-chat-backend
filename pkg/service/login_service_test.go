package service_test

import (
	"testing"

	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/service"
)

func TestLoginWorksAsExpected(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db)

	expectedID := 0
	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(username, password)

	login, err := service.Login(username, password)
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
	service := service.NewService(db)

	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(username, password)

	_, err := service.Login("user2", password)
	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotLoginInvalidPassword(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db)

	username := "user1"
	password := "generic_password"

	_, _ = service.CreateUser(username, password)

	_, err := service.Login(username, "invalidpwd")
	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}
