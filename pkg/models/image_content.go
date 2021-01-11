package models

import (
	"encoding/json"
	"fmt"
)

const ImageType string = "image"

// ImageContent represents an image with all its metadata.
// TODO: Extend with more metadata.
type ImageContent struct {
	URL string `json:"url"`
}

func (content ImageContent) Type() string {
	return ImageType
}

// Pass all the metadata in a string format with a separator to parse later.
func (content ImageContent) Data() string {
	return fmt.Sprintf("%s", content.URL)
}

// Get all the metadata from the string previously given by Data().
func (content *ImageContent) FromData(data string) {
	content.URL = data
	return
}

func (content ImageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		}{
			Type: ImageType,
			URL:  content.URL,
		},
	)
}
