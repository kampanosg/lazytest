package state

import (
	"time"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type HistoricalEntry struct {
	Timestamp time.Time
	Passed    bool
}

type State struct {
	TestTree    *tview.TreeNode
	TestOutput  map[*tview.TreeNode]*models.LazyTestResult
	FailedTests []*tview.TreeNode
	PassedTests []*tview.TreeNode
	IsSearching bool
	Size        *Size
	History     map[string][]string
}

type Size struct {
	Sidebar     int
	MainContent int
}

func NewState() *State {
	return &State{
		TestOutput:  make(map[*tview.TreeNode]*models.LazyTestResult),
		FailedTests: make([]*tview.TreeNode, 0),
		PassedTests: make([]*tview.TreeNode, 0),
		IsSearching: false,
		TestTree:    tview.NewTreeNode("."),
		History:     make(map[string][]string),
		Size: &Size{
			Sidebar:     4,
			MainContent: 8,
		},
	}
}

func (s *State) Reset() {
	s.FailedTests = make([]*tview.TreeNode, 0)
	s.PassedTests = make([]*tview.TreeNode, 0)
	// s.History = make(map[string][]string)
	s.Size = &Size{
		Sidebar:     4,
		MainContent: 8,
	}
}
