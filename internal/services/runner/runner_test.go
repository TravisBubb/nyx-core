package runner

import (
	"net/http"
	"testing"

	"github.com/TravisBubb/nyx-core/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRunnerBuildPostRequestSuccess(t *testing.T) {
	apiReq := &models.APIRequest{
		Method: "POST",
		Url:    "https://jsonplaceholder.typicode.com/posts",
		Body: map[string]interface{}{
			"prop1": "value1",
			"prop2": "value2",
		},
		Headers: map[string]string{
			"content-type": "application/json",
		},
	}

	runner := NewTestRunner()
	req, err := runner.buildRequest(apiReq)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, req.Method, http.MethodPost)
}

func TestRunnerExecuteTestCaseSuccess(t *testing.T) {
	testCase := &models.TestCase{
		Request: models.APIRequest{
			Method: "POST",
			Url:    "https://jsonplaceholder.typicode.com/posts",
			Body: map[string]interface{}{
				"prop1": "value1",
				"prop2": "value2",
			},
			Headers: map[string]string{
				"content-type": "application/json",
			},
		},
	}

	runner := NewTestRunner()
	resp, err := runner.execute(testCase)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRunnerExecuteTestSuiteSuccess(t *testing.T) {
	testSuite := &models.TestSuite{
		Version: "1.0",
		Tests: []models.TestCase{
			models.TestCase{
				Id:          "",
				Name:        "Create a new post",
				Description: "Tests the POST /posts endpoint",
				Request: models.APIRequest{
					Method: "POST",
					Url:    "https://jsonplaceholder.typicode.com/posts",
					Body: map[string]interface{}{
						"prop1": "value1",
						"prop2": "value2",
					},
					Headers: map[string]string{
						"content-type": "application/json",
					},
				},
			},
			models.TestCase{
				Id:          "",
				Name:        "Create a new post",
				Description: "Tests the POST /posts endpoint",
				Request: models.APIRequest{
					Method: "GET",
					Url:    "https://jsonplaceholder.typicode.com/todos/1",
					Headers: map[string]string{
						"content-type": "application/json",
					},
				},
			},
		},
	}

    runner := NewTestRunner()
    err := runner.Execute(testSuite)
    assert.NoError(t, err)
}
