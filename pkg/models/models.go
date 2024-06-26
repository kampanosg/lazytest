package models

import "time"

type LazyTestSuite struct {
	Tests []*LazyTest
	Path  string
}

type LazyTest struct {
	Name   string
	RunCmd string
}

type LazyTestResult struct {
	Passed   bool
	Output   string
	Duration time.Duration
}
