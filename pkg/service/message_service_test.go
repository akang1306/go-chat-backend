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
	expectedID := 0

	username1 := "user1"
	username2 := "user2"
	password := "b"
	user1, _ := service.CreateUser(ctx, username1, password)
	user2, _ := service.CreateUser(ctx, username2, password)

	msg, err := service.SendMessage(ctx, user1.ID, user2.ID, content)

	if err != nil {
		t.Errorf("Error: %s", err)
	} else {
		if msg.ID != expectedID {
			t.Errorf("Error: got id %d but expected %d", msg.ID, expectedID)
		}
		if msg.SenderID != user1.ID {
			t.Errorf("Error: got sender %d but expected %d", msg.SenderID, user1.ID)
		}
		if msg.RecipientID != user2.ID {
			t.Errorf("Error: got recipient %d but expected %d", msg.RecipientID, user2.ID)
		}
	}
}

func TestCanNotSendMessageToUnexistingRecipient(t *testing.T) {
	db := database.NewMockDB()
	service := service.NewService(db, jwt.New())
	ctx := context.TODO()

	username1 := "user1"
	password := "a"
	user1, _ := service.CreateUser(ctx, username1, password)
	content := models.StringContent{}

	_, err := service.SendMessage(ctx, user1.ID, user1.ID+1, content)

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
	user1, _ := service.CreateUser(ctx, username1, password)
	content := models.StringContent{}

	_, err := service.SendMessage(ctx, user1.ID+1, user1.ID, content)

	if err == nil {
		t.Errorf("Error: expected err but got nil")
	}
}
