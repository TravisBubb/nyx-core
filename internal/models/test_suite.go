package models

type TestSuite struct {
	Version string     `json:"version"`
	Tests   []TestCase `json:"tests"`
}
