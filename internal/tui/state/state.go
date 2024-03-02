package state

import (
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type State struct {
	TestTree    *tview.TreeNode
	TestOutput  map[*tview.TreeNode]*models.LazyTestResult
	FailedTests []*tview.TreeNode
	PassedTests []*tview.TreeNode
	IsSearching bool
}

func NewState() *State {
	return &State{
		TestOutput:  make(map[*tview.TreeNode]*models.LazyTestResult),
		FailedTests: make([]*tview.TreeNode, 0),
		PassedTests: make([]*tview.TreeNode, 0),
		IsSearching: false,
		TestTree:    tview.NewTreeNode("."),
	}
}

func (s *State) Reset() {
	s.FailedTests = make([]*tview.TreeNode, 0)
	s.PassedTests = make([]*tview.TreeNode, 0)
}
