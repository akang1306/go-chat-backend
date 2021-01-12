package service_test

import (
	"context"
	"testing"

	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/jwt"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/service"
)

func TestCanSendMessage(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	content := models.StringContent{}
	expectedID := models.MessageID(0)

	username1 := "user1"
	username2 := "user2"
	password := "b"
	userID1, _ := service.CreateUser(ctx, username1, password)
	userID2, _ := service.CreateUser(ctx, username2, password)

	msgInfo, err := service.SendMessage(ctx, userID1, userID2, content)

	if err != nil {
		t.Errorf("Error: %s", err)
	} else {
		if msgInfo.ID != expectedID {
			t.Errorf("Error: got id %d but expected %d", msgInfo.ID, expectedID)
		}
	}
}

func TestCanNotSendMessageToUnexistingRecipient(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username1 := "user1"
	password := "a"
	userID1, _ := service.CreateUser(ctx, username1, password)
	content := models.StringContent{}

	_, err := service.SendMessage(ctx, userID1, userID1+1, content)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotSendMessageFromUnexistingSender(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username1 := "user1"
	password := "a"
	userID1, _ := service.CreateUser(ctx, username1, password)
	content := models.StringContent{}

	_, err := service.SendMessage(ctx, userID1+1, userID1, content)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

func TestCanNotSendMessageToTheSameUser(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username1 := "user1"
	password := "a"
	userID1, _ := service.CreateUser(ctx, username1, password)
	content := models.StringContent{}

	_, err := service.SendMessage(ctx, userID1, userID1, content)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}

// TODO: Test all variations
func TestGetMessages(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username1 := "user1"
	username2 := "user2"
	password := "a"
	userID1, _ := service.CreateUser(ctx, username1, password)
	userID2, _ := service.CreateUser(ctx, username2, password)
	content := models.StringContent{}

	service.SendMessage(ctx, userID1, userID2, content)
	service.SendMessage(ctx, userID1, userID2, content)
	service.SendMessage(ctx, userID1, userID2, content)

	// limit works as expected
	limit := 2
	messages, err := service.GetMessages(ctx, userID2, 0, limit)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(messages) != limit {
		t.Errorf("Error: expected msg length %d but got %d", limit, len(messages))
	}
}
