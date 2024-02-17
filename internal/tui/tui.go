package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/kampanosg/lazytest/pkg/tree"
	"github.com/rivo/tview"
)

type runner interface {
	Run(command string) (*models.LazyTestResult, error)
}

type TUI struct {
	app          *tview.Application
	tree         *tview.TreeView
	output       *tview.TextView
	details      *tview.TextView
	legend       *tview.TextView
	flex         *tview.Flex
	state        state
	runner       runner
	lazyTestRoot *tree.LazyNode
}

func NewTUI(lt *tree.LazyNode, r runner) *TUI {
	return &TUI{
		app:          tview.NewApplication(),
		tree:         tview.NewTreeView(),
		output:       tview.NewTextView(),
		details:      tview.NewTextView(),
		legend:       tview.NewTextView(),
		flex:         tview.NewFlex(),
		state:        NewState(),
		runner:       r,
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
	t.setupDetails()
	t.setupLegend()
	t.setupFlex()

	t.app.EnableMouse(true)
	t.app.SetInputCapture(t.inputCapture)

	if err := t.app.SetRoot(t.flex, true).SetFocus(t.tree).Run(); err != nil {
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
		node.SetColor(tcell.ColorBlueViolet)
	})
	t.tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})
}

func (t *TUI) setupOutput() {
	t.output.SetBorder(true)
	t.output.SetTitle("Output")
	t.output.SetTitleAlign(tview.AlignLeft)
	t.output.SetBackgroundColor(tcell.ColorDefault)
	t.output.SetScrollable(true)
	t.output.SetDynamicColors(true)
}

func (t *TUI) setupDetails() {
	t.details.SetBorder(true)
	t.details.SetTitle("Details")
	t.details.SetTitleAlign(tview.AlignLeft)
	t.details.SetBackgroundColor(tcell.ColorDefault)
	t.details.SetScrollable(true)
	t.details.SetWrap(true)
	t.details.SetDynamicColors(true)
	t.details.SetText(fmt.Sprintf("[darkturquoise]Total: %d\n[limegreen]Passed: %d\n[indianred]Failed: %d",
		t.state.Details.TotalTests, t.state.Details.TotalPassed, t.state.Details.TotalFailed))
}

func (t *TUI) setupLegend() {
	t.legend.SetBorder(false)
	t.legend.SetTitleAlign(tview.AlignCenter)
	t.legend.SetBackgroundColor(tcell.ColorDefault)
	t.legend.SetText("?: help, 1/2/3: navigate, q: quit")
}

func (t *TUI) setupFlex() {
	sidebar := tview.NewFlex()
	sidebar.SetDirection(tview.FlexRow)

	sidebar.AddItem(t.tree, 0, 7, true)
	sidebar.AddItem(t.details, 0, 1, false)

	mainContent := tview.NewFlex()
	mainContent.AddItem(sidebar, 0, 1, false)
	mainContent.AddItem(t.output, 0, 2, false)

	footer := tview.NewFlex()
	footer.AddItem(t.legend, 0, 1, false)

	t.flex.SetDirection(tview.FlexRow)
	t.flex.AddItem(mainContent, 0, 30, false)
	t.flex.AddItem(footer, 0, 1, false)
}

func (t *TUI) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch pressed_key := event.Rune(); pressed_key {
	case 'q':
		t.app.Stop()
	case '1':
		t.app.SetFocus(t.tree)
	case '2':
		t.app.SetFocus(t.output)
	case '3':
		t.app.SetFocus(t.details)
	case 'r':
		t.handleRunCmd()
	}
	return event
}

func (t *TUI) buildTestNodes(lazyNode *tree.LazyNode) []*tview.TreeNode {
	nodes := []*tview.TreeNode{}

	if lazyNode.IsFolder && lazyNode.HasTestSuite() {
		f := tview.NewTreeNode(fmt.Sprintf("[default] %s", lazyNode.Name))
		f.SetSelectable(true)

		for _, child := range lazyNode.Children {
			ns := t.buildTestNodes(child)
			for _, n := range ns {
				f.AddChild(n)
			}
		}

		nodes = append(nodes, f)
	} else if !lazyNode.IsFolder {
		testSuite := tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", getNerdIcon(lazyNode.Suite.Type), lazyNode.Name))
		testSuite.SetSelectable(true)
		testSuite.SetReference(lazyNode.Suite)

		for _, t := range lazyNode.Suite.Tests {
			test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
			test.SetSelectable(true)
			test.SetReference(&t)
			testSuite.AddChild(test)
		}

		totalTests := t.state.Details.TotalTests + len(lazyNode.Suite.Tests)
		t.state.Details.TotalTests = totalTests

		nodes = append(nodes, testSuite)
	}
	return nodes
}

func (t *TUI) handleRunCmd() {
	testNode := t.tree.GetCurrentNode()
	if testNode == nil {
		return
	}

	ref := testNode.GetReference()
	if ref == nil {
		return
	}

	switch ref.(type) {
	case *models.LazyTestSuite:
		t.output.SetText("running suite" + ref.(*models.LazyTestSuite).Path)
	case *models.LazyTest:
		t.state.Details.TotalPassed = 0
		t.state.Details.TotalFailed = 0
		test := ref.(*models.LazyTest)
		testNode.SetText(fmt.Sprintf("[yellow] [darkturquoise]%s", test.Name))
		res, err := t.runner.Run(test.RunCmd)
		if err != nil {
			t.output.SetText(fmt.Sprintf("%s", err))
			t.output.SetBorderColor(tcell.ColorOrangeRed)
			testNode.SetText(fmt.Sprintf("[orangered] [darkturquoise]%s", test.Name))
			t.state.Details.TotalFailed++
		} else {
			t.output.SetText(res.Output)
			t.output.SetBorderColor(tcell.ColorGreen)
			testNode.SetText(fmt.Sprintf("[limegreen] [darkturquoise]%s", test.Name))
			t.state.Details.TotalPassed++
		}

		t.setupDetails()
	}
}

func getNerdIcon(suiteType string) string {
	switch suiteType {
	case "golang":
		return "󰟓"
	default:
		return ""
	}
}
