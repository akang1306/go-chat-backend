package database

import (
	"context"
	"errors"
	"sync"

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
	return nil, errors.New("user id not present")
}

func (db *MockDB) AddMessage(ctx context.Context, msg models.Message) error {
	db.m.Lock()
	defer db.m.Unlock()
	msg.ID = len(db.Messages)
	db.Messages = append(db.Messages, msg)
	return nil
}

func (db *MockDB) GetMessages(ctx context.Context, sender, start, limit int) error {
	return nil
}

func NewMockDB() *MockDB {
	return &MockDB{
		m:        sync.Mutex{},
		Users:    make([]models.User, 0),
		Messages: make([]models.Message, 0),
	}
}
