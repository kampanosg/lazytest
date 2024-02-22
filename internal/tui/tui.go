package tui

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/kampanosg/lazytest/pkg/tree"
	"github.com/rivo/tview"
)

const helpText = `
	[darkturquoise]1 / 2: [white]Focus on the tree / output 
	[darkturquoise]r: [white]Run the selected test / test suite
	[darkturquoise]a: [white]Run all tests
	[darkturquoise]f: [white]Run all failed tests
	[darkturquoise]p: [white]Run all passed tests
	[darkturquoise]/: [white]Search
	[darkturquoise]Enter: [white](in search mode) Go to the search results
	[darkturquoise]<ESC>: [white]Exit search mode
	[darkturquoise]C: [white](outside search mode) Clear search
	[darkturquoise]q: [white]Quit
	[darkturquoise]?: [white]Show this help message
`

type runner interface {
	Run(command string) *models.LazyTestResult
}

type TUI struct {
	app          *tview.Application
	tree         *tview.TreeView
	output       *tview.TextView
	infoBox      *tview.TextView
	search       *tview.InputField
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
		infoBox:      tview.NewTextView(),
		search:       tview.NewInputField(),
		legend:       tview.NewTextView(),
		flex:         tview.NewFlex(),
		state:        NewState(),
		runner:       r,
		lazyTestRoot: lt,
	}
}

func (t *TUI) Run() error {
	nodes := t.buildTestNodes(t.lazyTestRoot, "")
	var treeViewRoot *tview.TreeNode
	for i, n := range nodes {
		if i == 0 {
			treeViewRoot = n
			continue
		}
		treeViewRoot.AddChild(n)
	}

	t.state.Root = treeViewRoot

	t.setupTree(t.state.Root)
	t.setupOutput()
	t.setupInfoBox()
	t.setupSearch()
	t.setupLegend()
	t.setupFlex()

	t.app.EnableMouse(true)
	t.app.SetInputCapture(t.inputCapture)

	if err := t.app.SetRoot(t.flex, true).SetFocus(t.tree).EnablePaste(true).Run(); err != nil {
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
	t.tree.SetChangedFunc(t.nodeChanged)
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
	t.output.SetRegions(true)
}

func (t *TUI) setupInfoBox() {
	t.infoBox.SetBorder(true)
	t.infoBox.SetTitle("Info")
	t.infoBox.SetTitleAlign(tview.AlignLeft)
	t.infoBox.SetBackgroundColor(tcell.ColorDefault)
	t.infoBox.SetDynamicColors(true)
	t.infoBox.SetText("Welcome to LazyTest  ")
}

func (t *TUI) setupSearch() {
	t.search.SetTitle("Search")
	t.search.SetBorder(true)
	t.search.SetBackgroundColor(tcell.ColorDefault)
	t.search.SetTitleAlign(tview.AlignLeft)
	t.search.SetFieldBackgroundColor(tcell.ColorDefault)
	t.search.SetPlaceholder("Press / to search")
	t.search.SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorGray))

	t.search.SetFocusFunc(func() {
	})

	t.search.SetChangedFunc(func(searchQuery string) {
		if strings.HasSuffix(searchQuery, "/") {
			// when the user presses / to search, the / is still in the input field
			// so we're removing it here
			searchQuery = searchQuery[:len(searchQuery)-1]
			t.search.SetText(searchQuery)
		}

		if searchQuery != "" {
			nodes := t.buildTestNodes(t.lazyTestRoot, searchQuery)

			var root *tview.TreeNode
			for i, n := range nodes {
				if i == 0 {
					root = n
					continue
				}
				root.AddChild(n)
			}

			t.tree.SetRoot(root)
		} else {
			t.tree.SetRoot(t.state.Root)
		}
	})

	t.search.SetDoneFunc(func(key tcell.Key) {
		t.state.IsSearching = false

		if key == tcell.KeyEnter {
			t.app.SetFocus(t.tree)
		}

		if key == tcell.KeyEscape {
			t.search.SetText("")
			t.tree.SetRoot(t.state.Root)
			t.app.SetFocus(t.tree)
		}
	})
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
	sidebar.AddItem(t.tree, 0, 20, true)
	sidebar.AddItem(t.search, 0, 1, false)

	mainContent := tview.NewFlex()
	mainContent.SetDirection(tview.FlexRow)
	mainContent.AddItem(t.output, 0, 20, false)
	mainContent.AddItem(t.infoBox, 0, 1, false)

	app := tview.NewFlex()
	app.AddItem(sidebar, 0, 1, false)
	app.AddItem(mainContent, 0, 2, false)

	footer := tview.NewFlex()
	footer.AddItem(t.legend, 0, 1, false)

	t.flex.SetDirection(tview.FlexRow)
	t.flex.AddItem(app, 0, 30, false)
	t.flex.AddItem(footer, 0, 1, false)
}

func (t *TUI) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	if t.state.IsSearching {
		return event
	}
	switch pressed_key := event.Rune(); pressed_key {
	case 'q':
		t.app.Stop()
	case '1':
		t.app.SetFocus(t.tree)
	case '2':
		t.app.SetFocus(t.output)
	case 'r':
		go t.handleRunCmd()
	case 'a':
		go t.handleRunAllCmd()
	case 'f':
		go t.handleRunFailedCmd()
	case 'p':
		go t.handleRunPassedCmd()
	case '/':
		t.state.IsSearching = true
		t.app.SetFocus(t.search)
	case 'C':
		go t.handleClearSearchCmd()
	case '?':
		t.handleShowHelp()
	}
	return event
}

func (t *TUI) buildTestNodes(lazyNode *tree.LazyNode, searchQuery string) []*tview.TreeNode {
	nodes := []*tview.TreeNode{}

	if lazyNode.IsFolder && lazyNode.HasTestSuite() {
		f := tview.NewTreeNode(fmt.Sprintf("[default] %s", lazyNode.Name))
		f.SetSelectable(true)

		for _, child := range lazyNode.Children {
			ns := t.buildTestNodes(child, searchQuery)
			for _, n := range ns {
				f.AddChild(n)
			}
		}

		nodes = append(nodes, f)
	} else if !lazyNode.IsFolder {
		searchQuery = strings.ToLower(searchQuery)
		addTestSuite := false

		if searchQuery == "" {
			addTestSuite = true
		}

		testSuite := tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", getNerdIcon(lazyNode.Suite.Type), lazyNode.Name))
		testSuite.SetSelectable(true)
		testSuite.SetReference(lazyNode.Suite)

		if strings.Contains(strings.ToLower(lazyNode.Name), searchQuery) {
			addTestSuite = true
		}

		// can probably remove the i with Go 1.22
		for i, t := range lazyNode.Suite.Tests {
			test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
			test.SetSelectable(true)
			test.SetReference(&lazyNode.Suite.Tests[i])

			if searchQuery == "" || addTestSuite || strings.Contains(strings.ToLower(t.Name), searchQuery) {
				testSuite.AddChild(test)
				addTestSuite = true
			}
		}

		if addTestSuite {
			nodes = append(nodes, testSuite)
		}
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

	var wg sync.WaitGroup
	t.state.Reset()

	t.app.QueueUpdateDraw(func() {
		t.output.SetText("")
		t.infoBox.SetText("Running...")
	})

	switch ref.(type) {
	case *models.LazyTestSuite:
		for _, child := range testNode.GetChildren() {
			wg.Add(1)
			test := child.GetReference().(*models.LazyTest)
			go t.runTest(&wg, child, test)
		}
		wg.Wait()
		t.nodeChanged(testNode)
	case *models.LazyTest:
		wg.Add(1)
		t.runTest(&wg, testNode, ref.(*models.LazyTest))
		wg.Wait()
	}

	t.updateRunInfo()
}

func (t *TUI) handleRunAllCmd() {
	var wg sync.WaitGroup
	t.state.Reset()

	t.app.QueueUpdateDraw(func() {
		t.output.SetText("")
		t.infoBox.SetText("Running all tests...")
	})

	t.doRunAll(&wg, t.tree.GetRoot().GetChildren())

	wg.Wait()
	t.updateRunInfo()
}

func (t *TUI) doRunAll(wg *sync.WaitGroup, nodes []*tview.TreeNode) {
	for _, testNode := range nodes {
		if len(testNode.GetChildren()) > 0 {
			t.doRunAll(wg, testNode.GetChildren())
		} else {
			ref := testNode.GetReference()
			if ref == nil {
				continue
			}

			switch ref.(type) {
			case *models.LazyTest:
				wg.Add(1)
				t.runTest(wg, testNode, ref.(*models.LazyTest))
			}
		}
	}
}

func (t *TUI) handleRunFailedCmd() {
	if len(t.state.FailedTests) == 0 {
		t.app.QueueUpdateDraw(func() {
			t.infoBox.SetText("No failed tests to run. Good job ")
		})
		return
	}

	var wg sync.WaitGroup

	failedTests := t.state.FailedTests
	t.state.Reset()

	t.app.QueueUpdateDraw(func() {
		t.output.SetText("")
		t.infoBox.SetText("Running failed tests...")
	})

	for _, testNode := range failedTests {
		wg.Add(1)
		ref := testNode.GetReference()
		if ref == nil {
			continue
		}

		if test, ok := ref.(*models.LazyTest); ok {
			t.runTest(&wg, testNode, test)
		}
	}

	wg.Wait()
	t.updateRunInfo()
}

func (t *TUI) handleRunPassedCmd() {
	if len(t.state.PassedTests) == 0 {
		t.app.QueueUpdateDraw(func() {
			t.infoBox.SetText("No passed tests to run. Try running all tests ")
		})
		return
	}

	var wg sync.WaitGroup

	passedTests := t.state.PassedTests
	t.state.Reset()

	t.app.QueueUpdateDraw(func() {
		t.output.SetText("")
		t.infoBox.SetText("Running passed tests...")
	})

	for _, testNode := range passedTests {
		wg.Add(1)
		ref := testNode.GetReference()
		if ref == nil {
			continue
		}

		if test, ok := ref.(*models.LazyTest); ok {
			t.runTest(&wg, testNode, test)
		}

	}

	wg.Wait()
	t.updateRunInfo()
}

func (t *TUI) runTest(wg *sync.WaitGroup, testNode *tview.TreeNode, test *models.LazyTest) {
	defer wg.Done()

	t.app.QueueUpdateDraw(func() {
		testNode.SetText(fmt.Sprintf("[yellow] [darkturquoise]%s", test.Name))
		t.output.SetBorderColor(tcell.ColorYellow)
	})

	res := t.runner.Run(test.RunCmd)
	if res.IsSuccess {
		t.app.QueueUpdateDraw(func() {
			t.output.SetBorderColor(tcell.ColorGreen)
			testNode.SetText(fmt.Sprintf("[limegreen] [darkturquoise]%s", test.Name))
		})
		t.state.PassedTests = append(t.state.PassedTests, testNode)
	} else {
		t.app.QueueUpdateDraw(func() {
			t.output.SetBorderColor(tcell.ColorOrangeRed)
			testNode.SetText(fmt.Sprintf("[orangered] [darkturquoise]%s", test.Name))
		})
		t.state.FailedTests = append(t.state.FailedTests, testNode)
	}

	t.app.QueueUpdateDraw(func() {
		t.output.SetText(res.Output)
	})
	t.state.TestOutput[testNode] = res
}

func (t *TUI) handleShowHelp() {
	modal := tview.NewModal()
	modal.SetText(helpText)
	modal.SetBackgroundColor(tcell.ColorBlack)
	modal.AddButtons([]string{"Exit <ESC>"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex <= 1 {
			t.app.SetRoot(t.flex, true).SetFocus(t.tree)
		}
	})

	t.app.SetRoot(modal, true)
}

func (t *TUI) nodeChanged(node *tview.TreeNode) {
	node.SetColor(tcell.ColorBlueViolet)

	ref := node.GetReference()
	if ref == nil {
		return
	}

	switch ref.(type) {
	case *models.LazyTestSuite:
		borderColor := tcell.ColorWhite
		output := ""
		for _, child := range node.GetChildren() {
			res, ok := t.state.TestOutput[child]
			if ok {
				borderColor = tcell.ColorGreen
				if !res.IsSuccess {
					borderColor = tcell.ColorOrangeRed
				}

				o := fmt.Sprintf("--- %s ---\n%s\n\n", child.GetText(), res.Output)
				output = output + o
			}
		}
		t.output.SetBorderColor(borderColor)
		t.output.SetText(output)
	case *models.LazyTest:
		res, ok := t.state.TestOutput[node]
		if ok {
			if res.IsSuccess {
				t.output.SetBorderColor(tcell.ColorGreen)
			} else {
				t.output.SetBorderColor(tcell.ColorOrangeRed)
			}
			t.output.SetText(res.Output)
		}
	}
}

func (t *TUI) updateRunInfo() {
	t.app.QueueUpdateDraw(func() {
		totalFailed := len(t.state.FailedTests)
		totalPassed := len(t.state.PassedTests)
		msg := "Finished running."
		if totalPassed > 0 {
			msg = fmt.Sprintf("%s [limegreen]%d passed", msg, totalPassed)
		}
		if totalFailed > 0 {
			msg = fmt.Sprintf("%s. [orangered]%d failed", msg, totalFailed)
		}

		t.infoBox.SetText(msg)
	})
}

func (t *TUI) handleClearSearchCmd() {
	t.app.QueueUpdateDraw(func() {
		t.search.SetText("")
		t.tree.SetRoot(t.state.Root)
		t.app.SetFocus(t.tree)
	})
}

func getNerdIcon(suiteType string) string {
	switch suiteType {
	case "golang":
		return "󰟓"
	default:
		return ""
	}
}
