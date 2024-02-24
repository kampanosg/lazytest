package elements

import (
	"github.com/rivo/tview"
)

type elementData struct {
	TestTree *tview.TreeNode
}

type Elements struct {
	Flex      *tview.Flex
	Tree      *tview.TreeView
	Output    *tview.TextView
	Search    *tview.InputField
	InfoBox   *tview.TextView
	Legend    *tview.TextView
	HelpModal *tview.Modal
	data      *elementData
}

func NewElements(t *tview.TreeNode) *Elements {
	return &Elements{
		Flex:      tview.NewFlex(),
		Tree:      tview.NewTreeView(),
		Output:    tview.NewTextView(),
		InfoBox:   tview.NewTextView(),
		Search:    tview.NewInputField(),
		Legend:    tview.NewTextView(),
		HelpModal: tview.NewModal(),
		data: &elementData{
			TestTree: t,
		},
	}
}

func (e *Elements) Setup() {
	e.initTree()
	e.initOutput()
	e.initInfoBox()
	e.initSearch()
	e.initLegend()
	e.initHelp()
	e.initFlex()
}
