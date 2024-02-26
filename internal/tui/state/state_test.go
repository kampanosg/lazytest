package state

import (
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestState_Reset(t *testing.T) {
	type fields struct {
		TestTree    *tview.TreeNode
		TestOutput  map[*tview.TreeNode]*models.LazyTestResult
		FailedTests []*tview.TreeNode
		PassedTests []*tview.TreeNode
		IsSearching bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "resets state correctly",
			fields: fields{
				TestTree:    tview.NewTreeNode(""),
				TestOutput:  make(map[*tview.TreeNode]*models.LazyTestResult),
				FailedTests: []*tview.TreeNode{tview.NewTreeNode("")},
				PassedTests: []*tview.TreeNode{tview.NewTreeNode("")},
				IsSearching: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := tt

			t.Parallel()

			s := &State{
				TestTree:    tc.fields.TestTree,
				TestOutput:  tc.fields.TestOutput,
				FailedTests: tc.fields.FailedTests,
				PassedTests: tc.fields.PassedTests,
				IsSearching: tc.fields.IsSearching,
			}
			s.Reset()

			assert.Equal(t, len(s.FailedTests), 0)
			assert.Equal(t, len(s.PassedTests), 0)

		})
	}
}
