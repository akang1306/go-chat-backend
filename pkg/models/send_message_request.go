package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Provides the request structure for the SendMessage endpoint.
type SendMessageRequest struct {
	Sender    UserID         `json:"sender"`
	Recipient UserID         `json:"recipient"`
	Content   MessageContent `json:"content"`
}

// Used by UnmarshalJSON & test purposes.
type SendMessageRequestJSON struct {
	Sender    UserID                 `json:"sender"`
	Recipient UserID                 `json:"recipient"`
	Content   map[string]interface{} `json:"content"`
}

// UnmarshalJSON lets the message request dinamically parse the content data.
func (request *SendMessageRequest) UnmarshalJSON(b []byte) error {
	var requestJSON SendMessageRequestJSON
	if err := json.Unmarshal(b, &requestJSON); err != nil {
		return err
	}
	request.Sender = requestJSON.Sender
	request.Recipient = requestJSON.Recipient

	contentBytes, err := json.Marshal(requestJSON.Content)
	if err != nil {
		return err
	}

	contentType, ok := requestJSON.Content["type"].(string)
	if !ok {
		return errors.New("invalid content")
	}

	switch contentType {
	case "string":
		content := StringContent{}
		if err := json.Unmarshal(contentBytes, &content); err != nil {
			return err
		}
		request.Content = content
	case "image":
		content := ImageContent{}
		if err := json.Unmarshal(contentBytes, &content); err != nil {
			return err
		}
		request.Content = content
	default:
		return errors.New(fmt.Sprintf("unsupported type %s", contentType))
	}
	return nil
}
