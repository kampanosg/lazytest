package models

import "time"

type LazyTestSuite struct {
	Tests []*LazyTest
	Type  string
	Path  string
	Icon  string
}

type LazyTest struct {
	Name   string
	RunCmd string
}

type LazyTestResult struct {
	IsSuccess bool
	Output    string
	Duration time.Duration
}
