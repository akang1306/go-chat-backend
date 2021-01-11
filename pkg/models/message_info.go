package models

import "github.com/challenge/pkg/datetime"

type MessageInfo struct {
	ID        MessageID     `json:"id"`
	Timestamp datetime.Time `json:"timestamp"`
}
