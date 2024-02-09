package models

type LazyTestSuite interface {
	Tests() []LazyTest
	Path() string
}

type LazyTest interface {
	Name() string
	Run() (LazyTestResult, error)
}

type LazyTestResult interface {
	IsSuccess() bool
	Message() string
}
