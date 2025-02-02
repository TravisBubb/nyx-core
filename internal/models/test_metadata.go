package models

type TestMetadata struct {
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}
