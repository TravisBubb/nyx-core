package parser

import (
	"encoding/json"
	"testing"

	"github.com/TravisBubb/nyx-core/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestJsonParserSuccess(t *testing.T) {
	type testBody struct {
		Prop1 string `json:"prop1"`
		Prop2 string `json:"prop2"`
	}

	jsonStr := `
        {
            "version": "1.0",
            "tests": [
                {
                    "id": "",
                    "name": "Create a new post",
                    "description": "Tests the POST /posts endpoint",
                    "request": {
                        "url": "https://jsonplaceholder.typicode.com/posts",
                        "method": "POST",
                        "headers": {
                            "content-type": "application/json"
                        },
                        "body": {
                            "prop1": "value1",
                            "prop2": "value2"
                       }
                    }
                }
            ]
        }`

	expected := models.TestSuite{
		Version: "1.0",
		Tests: []models.TestCase{
			models.TestCase{
				Id:          "",
				Name:        "Create a new post",
				Description: "Tests the POST /posts endpoint",
				Request: models.APIRequest{
					Method: "POST",
					Url:    "https://jsonplaceholder.typicode.com/posts",
					Body: testBody{
						Prop1: "value1",
						Prop2: "value2",
					},
					Headers: map[string]string{
						"content-type": "application/json",
					},
				},
			},
		},
	}

	parser := NewJsonTestParser()
	testSuite, err := parser.Parse([]byte(jsonStr))

	assert.NoError(t, err)
	assert.NotNil(t, testSuite)

	assertTestSuitesEqual(t, &expected, testSuite)
}

// Assert that the two provided models.TestSuite objects are equal
func assertTestSuitesEqual(t *testing.T, expected, actual *models.TestSuite) {
	assert.Equal(t, actual.Version, expected.Version)

	assert.NotNil(t, actual.Tests)
	assert.Len(t, actual.Tests, 1)

	testCase := actual.Tests[0]
	expectedTestCase := expected.Tests[0]
	assert.Equal(t, testCase.Id, expectedTestCase.Id)
	assert.Equal(t, testCase.Name, expectedTestCase.Name)
	assert.Equal(t, testCase.Description, expectedTestCase.Description)

	apiRequest := testCase.Request
	expectedApiRequest := expectedTestCase.Request
	assert.Equal(t, apiRequest.Method, expectedApiRequest.Method)
	assert.Equal(t, apiRequest.Url, expectedApiRequest.Url)

	apiRequestBody, _ := json.Marshal(apiRequest.Body)
	expectedApiRequestBody, _ := json.Marshal(expectedApiRequest.Body)
	assert.Equal(t, apiRequestBody, expectedApiRequestBody)

	apiRequestHeaders, _ := json.Marshal(apiRequest.Headers)
	expectedApiRequestHeaders, _ := json.Marshal(expectedApiRequest.Headers)
	assert.Equal(t, apiRequestHeaders, expectedApiRequestHeaders)
}
