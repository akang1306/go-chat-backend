package models

import "encoding/json"

const StringType string = "string"

type StringContent struct {
	Value string `json:"text"`
}

func (content StringContent) Type() string {
	return StringType
}

func (content StringContent) Data() string {
	return content.Value
}

func (content *StringContent) FromData(data string) {
	content.Value = data
	return
}

func (content StringContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}{
			Type: content.Type(),
			Text: content.Value,
		},
	)
}
