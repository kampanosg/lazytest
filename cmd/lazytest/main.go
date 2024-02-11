package main

import (
	"log"

	"github.com/kampanosg/lazytest/internal/loader"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/tree"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	p2 := widgets.NewParagraph()
	p2.Text = "<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget"
	p2.Text = "test"
	p2.Title = "Output"

	legend := widgets.NewParagraph()
	legend.Text = "q: Quit"
	legend.Title = "Legend"

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	nodes := buildTestNodes(root)
	testsTree := widgets.NewTree()
	testsTree.TextStyle = ui.NewStyle(ui.ColorWhite)
	testsTree.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorCyan, ui.ModifierBold)
	testsTree.WrapText = false
	testsTree.SetNodes(nodes)
	testsTree.Title = "Tests"
	testsTree.ExpandAll()

	grid.Set(
		ui.NewRow(.90,
			ui.NewCol(1.0/3, testsTree),
			ui.NewCol(1.0/1.5, p2),
		),
		ui.NewRow(.1,
			ui.NewCol(1.0, legend),
		),
	)

	ui.Render(grid)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			testsTree.ScrollDown()
		case "k", "<Up>":
			testsTree.ScrollUp()
		case "<C-d>":
			testsTree.ScrollHalfPageDown()
		case "<C-u>":
			testsTree.ScrollHalfPageUp()
		case "<C-f>":
			testsTree.ScrollPageDown()
		case "<C-b>":
			testsTree.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				testsTree.ScrollTop()
			}
		case "<Home>":
			testsTree.ScrollTop()
		case "<Enter>":
			testsTree.ToggleExpand()
		case "G", "<End>":
			testsTree.ScrollBottom()
		case "E":
			testsTree.ExpandAll()
		case "C":
			testsTree.CollapseAll()
		case "<Resize>":
			x, y := ui.TerminalDimensions()
			testsTree.SetRect(0, 0, x, y)
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(testsTree)
	}

}

func buildTestNodes(n *tree.LazyNode) []*widgets.TreeNode {
	nodes := []*widgets.TreeNode{}
	if n.IsFolder && n.HasTestSuites() {
		f := &widgets.TreeNode{
			Value: nodeValue(n.Name),
			Nodes: []*widgets.TreeNode{},
		}
		for _, child := range n.Children {
			ns := buildTestNodes(child)
			f.Nodes = append(f.Nodes, ns...)
		}
		nodes = append(nodes, f)
	} else if !n.IsFolder {
		testSuite := &widgets.TreeNode{
			Value: nodeValue(n.Name),
			Nodes: []*widgets.TreeNode{},
		}

		for _, t := range n.Suite.Tests {
			test := &widgets.TreeNode{
				Value: nodeValue(t.Name),
			}
			testSuite.Nodes = append(testSuite.Nodes, test)
		}
		nodes = append(nodes, testSuite)
	}
	return nodes
}
