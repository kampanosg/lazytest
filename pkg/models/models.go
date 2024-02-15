package models

type LazyTestSuite struct {
	Tests []LazyTest
	Type  string
	Path  string
}

type LazyTest struct {
	Name string
	Run  func(LazyTestResult, error)
}

type LazyTestResult struct {
	IsSuccess bool
	Output    string
}
