package openapi

import "testing"

func TestV3ValidatorFailsYaml(t *testing.T) {
	input := `openapi: 3.0.0
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9

servers:
  - url: http://api.example.com/v1
    description: Optional server description, e.g. Main (production) server
  - url: http://staging-api.example.com
    description: Optional server description, e.g. Internal staging server for testing

paths:
  /users a malformed path
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML.
      responses:
        "200": # status code
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string`

	bytes := []byte(input)

	validator := &OpenAPIv3Validator{}
	err := validator.Validate(bytes)

	if err == nil {
		t.Fatal("Expected an error but validation succeeded")
	}
}

func TestV3ValidatorSucceedsYaml(t *testing.T) {

	input := `openapi: 3.0.0
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9

servers:
  - url: http://api.example.com/v1
    description: Optional server description, e.g. Main (production) server
  - url: http://staging-api.example.com
    description: Optional server description, e.g. Internal staging server for testing

paths:
  /users: 
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML.
      responses:
        "200": # status code
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string`

	bytes := []byte(input)

	validator := &OpenAPIv3Validator{}
	err := validator.Validate(bytes)

	if err != nil {
		t.Fatalf("Expected validation to succeed but validation failed with error %s", err)
	}
}
