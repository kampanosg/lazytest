package tui

import (
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type state struct {
	Details     details
	TestOutput  map[*tview.TreeNode]*models.LazyTestResult
	FailedTests []*tview.TreeNode
}

type details struct {
	TotalPassed int
}

func NewState() state {
	return state{
		Details: details{
			TotalPassed: 0,
		},
		TestOutput:  make(map[*tview.TreeNode]*models.LazyTestResult),
		FailedTests: make([]*tview.TreeNode, 0),
	}
}

func (s *state) Reset() {
	s.Details.TotalPassed = 0
	s.FailedTests = make([]*tview.TreeNode, 0)
}
