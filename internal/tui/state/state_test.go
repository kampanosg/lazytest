package state

import (
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	s := &State{
		TestTree:    tview.NewTreeNode(""),
		TestOutput:  make(map[*tview.TreeNode]*models.LazyTestResult),
		FailedTests: []*tview.TreeNode{tview.NewTreeNode("")},
		PassedTests: []*tview.TreeNode{tview.NewTreeNode("")},
		IsSearching: false,
	}

	s.Reset()

	assert.Empty(t, s.FailedTests)
	assert.Empty(t, s.FailedTests, 0)
	assert.Equal(t, 4, s.Size.Sidebar)
	assert.Equal(t, 8, s.Size.MainContent)
}
