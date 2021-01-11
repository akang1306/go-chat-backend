package models

import "encoding/json"

// ImageContent represents an image with all its metadata.
// TODO: Extend with more metadata.
type ImageContent struct {
	URL string `json:"url"`
}

func (content ImageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		}{
			Type: "image",
			URL:  content.URL,
		},
	)
}
