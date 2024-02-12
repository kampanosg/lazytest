package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/loader"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/tree"

	"github.com/rivo/tview"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func main() {
	currentDir := "."
	root := tree.NewFolder(currentDir)

	loader := loader.NewLazyTestLoader([]engines.LazyTestEngine{
		golang.NewGolangEngine(),
	})

	if err := loader.LoadLazyTests(currentDir, root); err != nil {
		panic(err)
	}

	app := tview.NewApplication()
	app.EnableMouse(true)

	tree := tview.NewTreeView()

	nodes := buildTestNodes(root)
	var treeViewRoot *tview.TreeNode
	for i, n := range nodes {
		if i == 0 {
			treeViewRoot = n
			continue
		}
		treeViewRoot.AddChild(n)
	}

	tree.SetTitle("Tests")
	tree.SetTitleAlign(tview.AlignLeft)
	tree.SetBorder(true)
	tree.SetBorderColor(tcell.ColorBlue)
	tree.SetRoot(treeViewRoot)
	tree.SetCurrentNode(treeViewRoot)
	tree.SetTopLevel(1)
	tree.SetBackgroundColor(tcell.ColorDefault)
	tree.SetChangedFunc(func(node *tview.TreeNode) {
		node.SetColor(tcell.ColorDarkSeaGreen)
	})
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})

	output := tview.NewTextView()
	output.SetTitle("Output")
	output.SetTitleAlign(tview.AlignLeft)
	output.SetBorder(true)
	output.SetBorderColor(tcell.ColorBlue)
	output.SetBackgroundColor(tcell.ColorDefault)

	legend := tview.NewTextView()
	legend.SetText("?: help, 1/2: navigate, q: quit")
	legend.SetTextAlign(tview.AlignCenter)
	legend.SetBackgroundColor(tcell.ColorDefault)

	grid := tview.NewGrid()
	grid.SetRows(0, 1)
	grid.SetColumns(33, 0)
	grid.SetBorders(false)
	grid.AddItem(tree, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(output, 0, 1, 1, 1, 0, 0, false)
	grid.AddItem(legend, 1, 0, 1, 2, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch pressed_key := event.Rune(); pressed_key {
		case 'q':
			app.Stop()
		case '1':
			app.SetFocus(tree)
		case '2':
			app.SetFocus(output)
		}
		return event
	})

	if err := app.SetRoot(grid, true).SetFocus(tree).Run(); err != nil {
		panic(err)
	}
}

func buildTestNodes(lazyNode *tree.LazyNode) []*tview.TreeNode {
	nodes := []*tview.TreeNode{}
	if lazyNode.IsFolder && lazyNode.HasTestSuites() {
		f := tview.NewTreeNode(fmt.Sprintf("[white] %s", lazyNode.Name))
		f.SetSelectable(true)

		for _, child := range lazyNode.Children {
			ns := buildTestNodes(child)
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
