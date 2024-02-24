package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Elements) initTree() {
	e.Tree.SetTitle("Tests")
	e.Tree.SetTitleAlign(tview.AlignLeft)
	e.Tree.SetBorder(true)
	e.Tree.SetBorderColor(tcell.ColorBlue)
	e.Tree.SetRoot(e.data.TestTree)
	e.Tree.SetCurrentNode(e.data.TestTree)
	e.Tree.SetTopLevel(0)
	e.Tree.SetBackgroundColor(tcell.ColorDefault)
	e.Tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})
}
