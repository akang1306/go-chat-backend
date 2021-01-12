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

// GetMessageContent parses the data with the matching content type struct.
// The format of the data is given by MessageContent.Data() and its parsing
// is given by MessageContent.FromData().
// This is used mainly to store and load the content dinamically from the database.
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
