package models

type APIRequest struct {
	Method      string            `json:"method"`
	Url         string            `json:"url"`
	Headers     map[string]string `json:"headers,omitempty"`
	QueryParams map[string]string `json:"query_params,omitempty"`
	Body        any               `json:"body,omitempty"`
}
