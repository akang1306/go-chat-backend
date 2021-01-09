package models

type Message struct {
	ID          int `json:"id"`
	SenderID    int `json:"sender"`
	RecipientID int `json:"recipient"`
	Content     MessageContent
}

type MessageContent interface{}
