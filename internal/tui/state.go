package tui

import (
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type state struct {
	Details    details
	TestOutput map[*tview.TreeNode]*models.LazyTestResult
}

type details struct {
	TotalTests  int
	TotalPassed int
	TotalFailed int
}

func NewState() state {
	return state{
		Details: details{
			TotalTests:  0,
			TotalPassed: 0,
			TotalFailed: 0,
		},
		TestOutput: make(map[*tview.TreeNode]*models.LazyTestResult),
	}
}
