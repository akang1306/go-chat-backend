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

func (db *MockDB) AddUser(ctx context.Context, username string, password []byte) (models.UserID, error) {
	for _, u := range db.Users {
		if u.Username == username {
			return 0, errors.New("Username already exists")
		}
	}
	db.m.Lock()
	defer db.m.Unlock()
	newID := len(db.Users)
	user := models.User{
		ID:       models.UserID(newID),
		Username: username,
		Password: password,
	}
	db.Users = append(db.Users, user)
	return user.ID, nil
}

func (db *MockDB) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	for _, user := range db.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user: %s not present", username))
}

func (db *MockDB) GetUserByID(ctx context.Context, id models.UserID) (*models.User, error) {
	for _, user := range db.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user id: %d not present", id))
}

func (db *MockDB) AddMessage(ctx context.Context, sender, recipient models.UserID, content models.MessageContent) (*models.MessageInfo, error) {
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
		ID:          models.MessageID(len(db.Messages)),
		Timestamp:   datetime.Now(),
		SenderID:    sender,
		RecipientID: recipient,
		Content:     content,
	}
	db.Messages = append(db.Messages, msg)
	return &models.MessageInfo{ID: msg.ID, Timestamp: msg.Timestamp}, nil
}

func (db *MockDB) GetMessages(ctx context.Context, sender models.UserID, start, limit int) ([]*models.Message, error) {
	msgList := make([]*models.Message, 0)
	for i, msg := range db.Messages {
		if int(msg.ID) >= start && msg.SenderID == sender && len(msgList) < limit {
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
