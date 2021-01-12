package database

import (
	"database/sql"
	"sync"
)

var db *sql.DB
var once sync.Once

// Commands for initial database creation.
const (
	UsersInitCmd = `CREATE TABLE IF NOT EXISTS users (
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL)`
	MessagesInitCmd = `CREATE TABLE IF NOT EXISTS messages (
		type TEXT NOT NULL,
		content TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		recipient TEXT NOT NULL,
		sender TEXT NOT NULL,
		FOREIGN KEY (recipient) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
		FOREIGN KEY (sender) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE)`
)

func InitChatDatabase() {
	db.Exec(UsersInitCmd)
	db.Exec(MessagesInitCmd)
}

func GetConnection(dbName string) *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite3", dbName)
		if err != nil {
			panic(err)
		}
	})
	return db
}

func NewConnection(dbName string) *sql.DB {
	return GetConnection(dbName)
}
