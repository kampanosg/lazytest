package elements

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func (e *Elements) initTree() {
	e.Tree.SetTitle("Tests")
	e.Tree.SetTitleAlign(tview.AlignLeft)
	e.Tree.SetBorder(true)
	e.Tree.SetBorderColor(tcell.ColorBlue)
	e.Tree.SetRoot(e.State.TestTree)
	e.Tree.SetCurrentNode(e.State.TestTree)
	e.Tree.SetTopLevel(0)
	e.Tree.SetBackgroundColor(tcell.ColorDefault)
	e.Tree.SetChangedFunc(e.HandleNodeChangedEvent)
	e.Tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})
}

func (e *Elements) HandleNodeChangedEvent(node *tview.TreeNode) {
	node.SetColor(tcell.ColorBlueViolet)

	ref := node.GetReference()
	if ref == nil {
		return
	}

	switch ref.(type) {
	case *models.LazyTestSuite:
		borderColor := tcell.ColorWhite
		outputs := ""
		for _, child := range node.GetChildren() {
			res, ok := e.State.TestOutput[child]
			if ok {
				borderColor = tcell.ColorGreen
				if !res.IsSuccess {
					borderColor = tcell.ColorOrangeRed
				}

				output := fmt.Sprintf("--- %s ---\n%s\n\n", child.GetText(), res.Output)
				outputs = outputs + output
			}
		}
		e.Output.SetBorderColor(borderColor)
		e.Output.SetText(outputs)
	case *models.LazyTest:
		res, ok := e.State.TestOutput[node]
		if ok {
			if res.IsSuccess {
				e.Output.SetBorderColor(tcell.ColorGreen)
			} else {
				e.Output.SetBorderColor(tcell.ColorOrangeRed)
			}
			e.Output.SetText(res.Output)
		}
	}
}
