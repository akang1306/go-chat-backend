package database

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/challenge/pkg/datetime"
	"github.com/challenge/pkg/models"
)

// MockDB works as an in memory database for testing.
type MockDB struct {
	m        sync.Mutex
	Users    []models.User
	Messages []models.Message
}

func (db *MockDB) AddUser(ctx context.Context, username string, password []byte) (*models.User, error) {
	for _, u := range db.Users {
		if u.Username == username {
			return nil, errors.New("Username already exists")
		}
	}
	db.m.Lock()
	defer db.m.Unlock()
	newID := len(db.Users)
	user := models.User{
		ID:       newID,
		Username: username,
		Password: password,
	}
	db.Users = append(db.Users, user)
	return &user, nil
}

func (db *MockDB) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	for _, user := range db.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user: %s not present", username))
}

func (db *MockDB) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	for _, user := range db.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user id: %d not present", id))
}

func (db *MockDB) AddMessage(ctx context.Context, sender, recipient int, content models.MessageContent) (*models.Message, error) {
	db.m.Lock()
	defer db.m.Unlock()

	// Check that sender & recipient both exist
	if _, err := db.GetUserByID(ctx, sender); err != nil {
		return nil, err
	}
	if _, err := db.GetUserByID(ctx, recipient); err != nil {
		return nil, err
	}

	msg := models.Message{
		ID:          len(db.Messages),
		Timestamp:   datetime.Now(),
		SenderID:    sender,
		RecipientID: recipient,
		Content:     content,
	}
	db.Messages = append(db.Messages, msg)
	return &msg, nil
}

func (db *MockDB) GetMessages(ctx context.Context, sender, start, limit int) ([]*models.Message, error) {
	msgList := make([]*models.Message, 0)
	for i, msg := range db.Messages {
		if msg.ID >= start && msg.SenderID == sender && len(msgList) < limit {
			msgList = append(msgList, &db.Messages[i])
		}
	}
	return msgList, nil
}

func NewMockDB() *MockDB {
	return &MockDB{
		m:        sync.Mutex{},
		Users:    make([]models.User, 0),
		Messages: make([]models.Message, 0),
	}
}
