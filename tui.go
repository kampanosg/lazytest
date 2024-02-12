package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/pkg/tree"
	"github.com/rivo/tview"
)

type TUI struct {
	app          *tview.Application
	tree         *tview.TreeView
	output       *tview.TextView
	legend       *tview.TextView
	grid         *tview.Grid
	lazyTestRoot *tree.LazyNode
}

func NewTUI(lt *tree.LazyNode) *TUI {
	return &TUI{
		app:          tview.NewApplication(),
		tree:         tview.NewTreeView(),
		output:       tview.NewTextView(),
		legend:       tview.NewTextView(),
		grid:         tview.NewGrid(),
		lazyTestRoot: lt,
	}
}

func (t *TUI) Run() error {
	nodes := t.buildTestNodes(t.lazyTestRoot)
	var treeViewRoot *tview.TreeNode
	for i, n := range nodes {
		if i == 0 {
			treeViewRoot = n
			continue
		}
		treeViewRoot.AddChild(n)
	}

	t.setupTree(treeViewRoot)
	t.setupOutput()
	t.setupLegend()
	t.setupGrid()

	t.app.EnableMouse(true)
	t.app.SetInputCapture(t.inputCapture)

	if err := t.app.SetRoot(t.grid, true).SetFocus(t.tree).Run(); err != nil {
		return err
	}

	return nil
}

func (t *TUI) setupTree(treeViewRoot *tview.TreeNode) {
	t.tree.SetTitle("Tests")
	t.tree.SetTitleAlign(tview.AlignLeft)
	t.tree.SetBorder(true)
	t.tree.SetBorderColor(tcell.ColorBlue)
	t.tree.SetRoot(treeViewRoot)
	t.tree.SetCurrentNode(treeViewRoot)
	t.tree.SetTopLevel(1)
	t.tree.SetBackgroundColor(tcell.ColorDefault)
	t.tree.SetChangedFunc(func(node *tview.TreeNode) {
		node.SetColor(tcell.ColorDarkSeaGreen)
	})
	t.tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})
}

func (t *TUI) setupOutput() {
	t.output.SetBorder(true)
	t.output.SetTitle("Output")
	t.output.SetTitleAlign(tview.AlignLeft)
	t.output.SetBorderColor(tcell.ColorBlue)
	t.output.SetBackgroundColor(tcell.ColorDefault)
}

func (t *TUI) setupLegend() {
	t.legend.SetBorder(true)
	t.legend.SetTitleAlign(tview.AlignCenter)
	t.legend.SetBackgroundColor(tcell.ColorDefault)
	t.legend.SetText("?: help, 1/2: navigate, q: quit")
}

func (t *TUI) setupGrid() {
	t.grid.SetRows(0, 1)
	t.grid.SetColumns(33, 0)
	t.grid.SetBorders(false)
	t.grid.AddItem(t.tree, 0, 0, 1, 1, 0, 0, true)
	t.grid.AddItem(t.output, 0, 1, 1, 1, 0, 0, false)
	t.grid.AddItem(t.legend, 1, 0, 1, 2, 0, 0, false)
}

func (t *TUI) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch pressed_key := event.Rune(); pressed_key {
	case 'q':
		t.app.Stop()
	case '1':
		t.app.SetFocus(t.tree)
	case '2':
		t.app.SetFocus(t.output)
	}
	return event
}

func (t *TUI) buildTestNodes(lazyNode *tree.LazyNode) []*tview.TreeNode {
	nodes := []*tview.TreeNode{}
	if lazyNode.IsFolder && lazyNode.HasTestSuites() {
		f := tview.NewTreeNode(fmt.Sprintf("[white] %s", lazyNode.Name))
		f.SetSelectable(true)

		for _, child := range lazyNode.Children {
			ns := t.buildTestNodes(child)
			for _, n := range ns {
				f.AddChild(n)
			}
		}

		nodes = append(nodes, f)
	} else if !lazyNode.IsFolder {
		testSuite := tview.NewTreeNode(fmt.Sprintf("[white]󰟓 %s", lazyNode.Name))
		testSuite.SetSelectable(true)

		for _, t := range lazyNode.Suite.Tests {
			test := tview.NewTreeNode(fmt.Sprintf("[blue] 󰐊 %s", t.Name))
			test.SetSelectable(true)
			testSuite.AddChild(test)
		}

		nodes = append(nodes, testSuite)
	}
	return nodes
}
