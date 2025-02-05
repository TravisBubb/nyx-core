package models

// Represents the results of executing a test suite
type TestSummary struct {
	RequestCount int
	TestsPassed  int
	TestsFailed  int
}
