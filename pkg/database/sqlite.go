package database

import (
	"context"
	"database/sql"

	"github.com/challenge/pkg/datetime"
	"github.com/challenge/pkg/models"
)

type SQLiteDB struct {
	conn *sql.DB
}

func (db *SQLiteDB) AddUser(ctx context.Context, username string, password []byte) (models.UserID, error) {
	res, err := db.conn.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *SQLiteDB) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	rows, err := db.conn.Query(
		`SELECT id, username, password FROM users
			WHERE username = ?`, username)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	rows.Next()
	defer rows.Close()
	if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *SQLiteDB) AddMessage(ctx context.Context, sender, recipient models.UserID, content models.MessageContent) (*models.MessageInfo, error) {
	timestamp := datetime.Now()
	res, err := db.conn.Exec("INSERT INTO messages (timestamp, type, content, recipient, sender) VALUES (?, ?, ?, ?, ?)",
		timestamp.String(), content.Type(), content.Data(), recipient, sender)
	if err != nil {
		return nil, err
	}
	id, idErr := res.LastInsertId()
	if idErr != nil {
		return nil, idErr
	}
	return &models.MessageInfo{
		ID:        id,
		Timestamp: timestamp,
	}, nil
}

func (db *SQLiteDB) GetMessages(ctx context.Context, recipient models.UserID, start, limit int) ([]*models.Message, error) {
	rows, err := db.conn.Query(
		`SELECT rowid, timestamp, type, content, recipient, sender FROM messages
			WHERE rowid >= ? AND recipient = ? ORDER BY rowid ASC LIMIT ?`,
		start, recipient, limit)
	if err != nil {
		return nil, err
	}
	_ = models.Message{}
	defer rows.Close()

	messages := make([]*models.Message, 0)
	for rows.Next() {
		var contentData, contentType, timestamp string
		msg := models.Message{}
		if err := rows.Scan(&msg.ID, &timestamp, &contentType, &contentData, &msg.RecipientID, &msg.SenderID); err != nil {
			return nil, err
		}
		time, timeErr := datetime.Parse(timestamp)
		if timeErr != nil {
			return nil, timeErr
		}
		msg.Timestamp = time
		msg.Content = models.GetMessageContent(contentType, contentData)
		messages = append(messages, &msg)
	}
	return messages, nil
}

func NewSQLiteDB(conn *sql.DB) *SQLiteDB {
	return &SQLiteDB{conn: conn}
}
