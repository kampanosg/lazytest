package elements

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	root := tview.NewTreeNode("root")
	root.SetReference("ref")

	e := NewElements()

	e.Setup(root, nil, nil, nil, nil)

	assert.Equal(t, "Tests", e.Tree.GetTitle())
	assert.Equal(t, tcell.ColorDefault, e.Tree.GetBackgroundColor())
	assert.Equal(t, root, e.Tree.GetRoot())
	assert.Equal(t, root, e.Tree.GetCurrentNode())
	assert.Equal(t, "ref", e.Tree.GetCurrentNode().GetReference())
}
