package models

import "github.com/challenge/pkg/datetime"

type MessageID = int64

type MessageContent interface {
	Type() string
	Data() string
}

type Message struct {
	ID          MessageID     `json:"id"`
	Timestamp   datetime.Time `json:"timestamp"`
	SenderID    UserID        `json:"sender"`
	RecipientID UserID        `json:"recipient"`
	Content     MessageContent
}

func GetMessageContent(contentType string, data string) MessageContent {
	switch contentType {
	case ImageType:
		c := &ImageContent{}
		c.FromData(data)
		return c
	}
	// Setting string as default
	c := &StringContent{}
	c.FromData(data)
	return c
}
