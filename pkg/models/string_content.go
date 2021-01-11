package models

import "encoding/json"

type StringContent struct {
	Value string `json:"text"`
}

func (content StringContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}{
			Type: "string",
			Text: content.Value,
		},
	)
}
