package models

type TestCase struct {
	Id          string       `json:"id"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
    Request     APIRequest   `json:"request"`
	Expected    APIResponse  `json:"expected,omitempty"`
	Metadata    TestMetadata `json:"metadata,omitempty"`
}
