package models

import "github.com/challenge/pkg/datetime"

type MessageContent interface{}

type Message struct {
	ID          int           `json:"id"`
	Timestamp   datetime.Time `json:"timestamp"`
	SenderID    int           `json:"sender"`
	RecipientID int           `json:"recipient"`
	Content     MessageContent
}
