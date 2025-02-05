package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TravisBubb/nyx-core/internal/models"
)

type TestRunner struct {
	client       *http.Client
	requestsSent int
	testsPassed  int
	testsFailed  int
}

func NewTestRunner() *TestRunner {
	return &TestRunner{
		client: &http.Client{},
	}
}

// Execute the provided TestSuite and return the results
func (r *TestRunner) Execute(s *models.TestSuite) (models.TestSummary, error) {
	if len(s.Tests) == 0 {
		return models.TestSummary{}, nil
	}

	for _, test := range s.Tests {
		_, _ = r.execute(&test)
	}

	return models.TestSummary{
		RequestCount: r.requestsSent,
		TestsPassed:  r.testsPassed,
		TestsFailed:  r.testsFailed,
	}, nil
}

// Execute the given TestCase
func (r *TestRunner) execute(t *models.TestCase) (*models.APIResponse, error) {
	req, err := r.buildRequest(&t.Request)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Sending request: %+v\n\n", req)
    r.requestsSent++

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp, err := r.buildResponse(resp)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Received response: %+v\n\n", apiResp)

	// TODO: Execute any configured assertions for the TestCase

	return apiResp, nil
}

// Build a net/http Request object from the given APIRequest
func (r *TestRunner) buildRequest(apiReq *models.APIRequest) (*http.Request, error) {
	var (
		reqBody []byte
		err     error
	)

	if apiReq.Body != nil {
		reqBody, err = json.Marshal(apiReq.Body)
	}

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(apiReq.Method, apiReq.Url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	for k, v := range apiReq.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

// Build an APIResponse object from the given net/http Response
func (r *TestRunner) buildResponse(resp *http.Response) (*models.APIResponse, error) {
	var apiResponse models.APIResponse
	apiResponse.StatusCode = resp.StatusCode

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &apiResponse.Body)
	if err != nil {
		return nil, err
	}

	apiResponse.Headers = map[string]string{}
	for k, v := range resp.Header {
		apiResponse.Headers[k] = v[0]
	}

	return &apiResponse, nil
}
