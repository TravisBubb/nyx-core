package models

type APIResponse struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       any               `json:"body,omitempty"`
	Assertions []Assertion       `json:"assertions,omitempty"`
}
