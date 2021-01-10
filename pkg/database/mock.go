package database

import (
	"errors"

	"github.com/challenge/pkg/models"
)

// MockDB works as an in memory database for testing.
type MockDB struct {
	Users    []models.User
	Messages []models.Message
}

func (db *MockDB) AddUser(user models.User) (int, error) {
	for _, u := range db.Users {
		if u.Username == user.Username {
			return 0, errors.New("Username already exists")
		}
	}
	newID := len(db.Users)
	user.ID = newID
	db.Users = append(db.Users, user)
	return newID, nil
}

func (db *MockDB) GetUserByUsername(username string) (*models.User, error) {
	for _, user := range db.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user id not present")
}

func (db *MockDB) AddMessage(msg models.Message) error {
	msg.ID = len(db.Messages)
	db.Messages = append(db.Messages, msg)
	return nil
}

func (db *MockDB) GetMessages(sender, start, limit int) error {
	return nil
}

func NewMockDB() *MockDB {
	return &MockDB{
		Users:    make([]models.User, 0),
		Messages: make([]models.Message, 0),
	}
}
