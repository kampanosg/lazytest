package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type elementData struct {
	TestTree *tview.TreeNode
}

type handlers struct {
	handleTreeChanged   func(node *tview.TreeNode)
	handleSearchDone    func(key tcell.Key)
	handleSearchChanged func(query string)
}

type Elements struct {
	Flex      *tview.Flex
	Tree      *tview.TreeView
	Output    *tview.TextView
	Search    *tview.InputField
	InfoBox   *tview.TextView
	Legend    *tview.TextView
	HelpModal *tview.Modal

	data     *elementData
	handlers *handlers
}

func NewElements(
	t *tview.TreeNode,
	htc func(node *tview.TreeNode),
	hsc func(query string),
	hsd func(key tcell.Key),
) *Elements {
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
		handlers: &handlers{
			handleTreeChanged:   htc,
			handleSearchChanged: hsc,
			handleSearchDone:    hsd,
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
