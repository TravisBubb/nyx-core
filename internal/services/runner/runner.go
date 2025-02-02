package runner

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/TravisBubb/nyx-core/internal/models"
)

type TestRunner struct {
	client *http.Client
}

func NewTestRunner() *TestRunner {
	return &TestRunner{
		client: &http.Client{},
	}
}

// Execute the provided TestSuite and return the results
func (r *TestRunner) Execute(s *models.TestSuite) error {
	if len(s.Tests) == 0 {
		return nil
	}

	for _, test := range s.Tests {
		_, err := r.execute(&test)
		if err != nil {
			return err
		}
	}

	return nil
}

// Execute the given TestCase
func (r *TestRunner) execute(t *models.TestCase) (*models.APIResponse, error) {
	req, err := r.buildRequest(&t.Request)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
    
    apiResp, err := r.buildResponse(resp)
	if err != nil {
		return nil, err
	}
		
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
