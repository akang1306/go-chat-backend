package models

import "github.com/challenge/pkg/datetime"

type MessageID = int64

type MessageContent interface{}

type Message struct {
	ID          MessageID     `json:"id"`
	Timestamp   datetime.Time `json:"timestamp"`
	SenderID    UserID        `json:"sender"`
	RecipientID UserID        `json:"recipient"`
	Content     MessageContent
}
