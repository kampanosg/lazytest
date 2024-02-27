package utils

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	internal := tview.NewTreeNode("internal")
	apprentices := tview.NewTreeNode("apprentices")

	employees := tview.NewTreeNode("employees")

	darthVader := tview.NewTreeNode("Darth Vader")
	darthSidious := tview.NewTreeNode("Darth Sidious")
	darthMaul := tview.NewTreeNode("Darth Maul")

	testTree := tview.NewTreeNode(".").
		AddChild(internal.
			AddChild(employees.
				AddChild(darthVader).
				AddChild(darthSidious),
			).
			AddChild(apprentices.
				AddChild(darthMaul),
			),
		)

	tests := []struct {
		name  string
		query string
		want  []*tview.TreeNode
	}{
		// {
		// 	name:  "empty query",
		// 	query: "",
		// 	want:  []*tview.TreeNode{},
		// },
		// {
		// 	name:  "no results",
		// 	query: "luke",
		// 	want:  nil,
		// },
		{
			name:  "match",
			query: "Vader",
			want: []*tview.TreeNode{
				darthVader,
			},
		},
	}
	for _, tt := range tests {
		tc := tt

		t.Run(tc.name, func(t *testing.T) {
			res := Search(testTree, tc.query)
			assert.ElementsMatch(t, tc.want, res.GetChildren())
		})
	}
}
