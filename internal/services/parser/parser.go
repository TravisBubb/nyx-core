package parser

import (
	"encoding/json"
	"os"

	"github.com/TravisBubb/nyx-core/internal/models"
)

type TestParser interface {
	ParseFile(filepath string) (*models.TestSuite, error)
	Parse(data []byte) (*models.TestSuite, error)
}

type JsonTestParser struct {}

func NewJsonTestParser() *JsonTestParser {
	return &JsonTestParser{}
}

func (p *JsonTestParser) ParseFile(filepath string) (*models.TestSuite, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

    return p.Parse(data)
}

func (p *JsonTestParser) Parse(data []byte) (*models.TestSuite, error) {
	var testSuite models.TestSuite
	if err := json.Unmarshal(data, &testSuite); err != nil {
		return nil, err
	}

	return &testSuite, nil
}
