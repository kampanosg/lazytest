package tui

import (
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type state struct {
	TestOutput  map[*tview.TreeNode]*models.LazyTestResult
	FailedTests []*tview.TreeNode
	PassedTests []*tview.TreeNode
	IsSearching bool
	Root        *tview.TreeNode
}

func NewState() state {
	return state{
		TestOutput:  make(map[*tview.TreeNode]*models.LazyTestResult),
		FailedTests: make([]*tview.TreeNode, 0),
		PassedTests: make([]*tview.TreeNode, 0),
		IsSearching: false,
		Root:        tview.NewTreeNode(""),
	}
}

func (s *state) Reset() {
	s.FailedTests = make([]*tview.TreeNode, 0)
	s.PassedTests = make([]*tview.TreeNode, 0)
}
