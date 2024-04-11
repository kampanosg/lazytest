package utils

import (
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	vaderTests := &models.LazyTestSuite{
		Path: "Darth Vader",
		Tests: []*models.LazyTest{
			{
				Name: "TestVader",
			},
		},
	}

	sidiusTests := &models.LazyTestSuite{
		Path: "Darth Sidius",
		Tests: []*models.LazyTest{
			{
				Name: "TestSidius",
			},
		},
	}

	maulTests := &models.LazyTestSuite{
		Path: "Darth Maul",
		Tests: []*models.LazyTest{
			{
				Name: "TestMaul",
			},
		},
	}

	internal := tview.NewTreeNode("internal")
	apprentices := tview.NewTreeNode("apprentices")

	employees := tview.NewTreeNode("employees")

	darthVader := tview.NewTreeNode("Darth Vader")
	darthSidious := tview.NewTreeNode("Darth Sidious")
	darthMaul := tview.NewTreeNode("Darth Maul")

	darthVader.SetReference(vaderTests)
	darthSidious.SetReference(sidiusTests)
	darthMaul.SetReference(maulTests)

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
		{
			name:  "no results",
			query: "luke",
			want:  nil,
		},
		{
			name:  "match",
			query: "Vader",
			want: []*tview.TreeNode{
				darthVader,
			},
		},
		{
			name:  "match multiple",
			query: "Darth",
			want: []*tview.TreeNode{
				darthVader,
				darthSidious,
				darthMaul,
			},
		},
		{
			name:  "strict uppercase",
			query: "darth",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Search(testTree, tt.query)
			assert.ElementsMatch(t, tt.want, res.GetChildren())
		})
	}
}
