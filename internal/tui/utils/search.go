package utils

import (
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func Search(root *tview.TreeNode, query string) *tview.TreeNode {
	filtered := tview.NewTreeNode("Search results")
	doSearch(root, filtered, query)
	return filtered
}

func doSearch(original, filtered *tview.TreeNode, query string) {
	if query == "" {
		filtered.AddChild(original)
		return
	}

	ref := original.GetReference()
	if ref == nil {
		for _, child := range original.GetChildren() {
			doSearch(child, filtered, query)
		}
	} else {
		if testSuite, ok := ref.(*models.LazyTestSuite); ok {
			if strings.Contains(testSuite.Path, query) {
				filtered.AddChild(original)
				return
			}

			for _, test := range original.GetChildren() {
				ref := test.GetReference()
				if ref == nil {
					continue
				}

				if t, ok := ref.(*models.LazyTest); ok {
					if strings.Contains(t.Name, query) {
						filtered.AddChild(test)
					}
				}
			}
		}
	}
}
