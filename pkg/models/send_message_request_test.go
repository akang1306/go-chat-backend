package models_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/challenge/pkg/models"
)

func TestSendMessageRequestUnmarshal(t *testing.T) {
	cases := []struct {
		name            string
		JSON            models.SendMessageRequestJSON
		expectedRequest models.SendMessageRequest
		err             error
	}{
		{
			name: "valid string unmarshal",
			JSON: models.SendMessageRequestJSON{
				Sender:    0,
				Recipient: 1,
				Content: map[string]interface{}{
					"type": "string",
					"text": "some text",
				},
			},
			expectedRequest: models.SendMessageRequest{
				Sender:    0,
				Recipient: 1,
				Content: models.StringContent{
					Value: "some text",
				},
			},
		}, {
			name: "valid image unmarshal",
			JSON: models.SendMessageRequestJSON{
				Sender:    0,
				Recipient: 1,
				Content: map[string]interface{}{
					"type": "image",
					"url":  "www.url.com",
				},
			},
			expectedRequest: models.SendMessageRequest{
				Sender:    0,
				Recipient: 1,
				Content: models.ImageContent{
					URL: "www.url.com",
				},
			},
		}, {
			name: "invalid type unmarshal",
			JSON: models.SendMessageRequestJSON{
				Sender:    0,
				Recipient: 1,
				Content: map[string]interface{}{
					"type": "invalid",
					"url":  "www.url.com",
				},
			},
			err: errors.New("error"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			json, _ := json.Marshal(c.JSON)

			request := &models.SendMessageRequest{}
			err := request.UnmarshalJSON([]byte(json))

			if err != nil {
				if c.err == nil {
					t.Errorf("%s: unexpected error (%s)", c.name, err.Error())
				}
			} else {
				contentType := reflect.TypeOf(request.Content)
				expectedContentType := reflect.TypeOf(c.expectedRequest.Content)

				if contentType != expectedContentType {
					t.Errorf("%s: expected type (%v) but got (%v)", c.name, expectedContentType, contentType)
				}

				if !reflect.DeepEqual(c.expectedRequest, *request) {
					t.Errorf("%s: expected value (%v) but got (%v)", c.name, c.expectedRequest, *request)
				}
			}
		})
	}
}
