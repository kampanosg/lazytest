package tui

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/loader"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
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
	app       *tview.Application
	Elements  *elements.Elements
	tree      *tview.TreeView
	search    *tview.InputField
	flex      *tview.Flex
	state     state
	directory string
	runner    runner
	loader    *loader.LazyTestLoader
}

func NewTUI(d string, r runner, e []engines.LazyEngine) *TUI {
	return &TUI{
		app:       tview.NewApplication(),
		Elements:  elements.NewElements(),
		tree:      tview.NewTreeView(),
		search:    tview.NewInputField(),
		flex:      tview.NewFlex(),
		state:     NewState(),
		directory: d,
		runner:    r,
		loader:    loader.NewLazyTestLoader(e),
	}
}

func (t *TUI) Run() error {
	t.state.Root = tview.NewTreeNode(t.directory)
	t.loader.LoadLazyTests(t.directory, t.state.Root)

	t.Elements.Setup()

	t.setupTree(t.state.Root)
	t.setupSearch()
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
	t.tree.SetTopLevel(0)
	t.tree.SetBackgroundColor(tcell.ColorDefault)
	t.tree.SetChangedFunc(t.nodeChanged)
	t.tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})
}

func (t *TUI) setupSearch() {
	t.search.SetTitle("Search")
	t.search.SetBorder(true)
	t.search.SetBackgroundColor(tcell.ColorDefault)
	t.search.SetTitleAlign(tview.AlignLeft)
	t.search.SetFieldBackgroundColor(tcell.ColorDefault)
	t.search.SetPlaceholder("Press / to search")
	t.search.SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorGray))

	t.search.SetChangedFunc(func(searchQuery string) {
		if strings.HasSuffix(searchQuery, "/") {
			// when the user presses / to search, the / is still in the input field
			// so we're removing it here
			searchQuery = searchQuery[:len(searchQuery)-1]
			t.search.SetText(searchQuery)
		}

		if searchQuery == "" {
			t.tree.SetRoot(t.state.Root)
			return
		}

		root := t.state.Root
		filtered := search(root, searchQuery)
		t.tree.SetRoot(filtered)

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
			t.Elements.InfoBox.SetText("Exited search mode")
		}
	})
}

func search(root *tview.TreeNode, query string) *tview.TreeNode {
	filtered := tview.NewTreeNode("Search results")
	doSearch(root, filtered, query)
	return filtered
}

func doSearch(original, filtered *tview.TreeNode, query string) {
	if query == "" {
		filtered.AddChild(original)
		return
	}

	ref := original.GetReference()
	if ref == nil {
		for _, child := range original.GetChildren() {
			doSearch(child, filtered, query)
		}
	} else {
		if testSuite, ok := ref.(*models.LazyTestSuite); ok {
			if strings.Contains(testSuite.Path, query) {
				filtered.AddChild(original)
				return
			}

			for _, test := range original.GetChildren() {
				ref := test.GetReference()
				if ref == nil {
					continue
				}

				if t, ok := ref.(*models.LazyTest); ok {
					if strings.Contains(t.Name, query) {
						filtered.AddChild(test)
					}
				}
			}
		}
	}
}

func (t *TUI) setupFlex() {
	sidebar := tview.NewFlex()
	sidebar.SetDirection(tview.FlexRow)
	sidebar.AddItem(t.tree, 0, 20, true)
	sidebar.AddItem(t.search, 3, 0, false)

	mainContent := tview.NewFlex()
	mainContent.SetDirection(tview.FlexRow)
	mainContent.AddItem(t.Elements.Output, 0, 20, false)
	mainContent.AddItem(t.Elements.InfoBox, 3, 0, false)

	app := tview.NewFlex()
	app.AddItem(sidebar, 0, 1, false)
	app.AddItem(mainContent, 0, 2, false)

	footer := tview.NewFlex()
	footer.AddItem(t.Elements.Legend, 0, 1, false)

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
		t.app.SetFocus(t.Elements.Output)
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
		t.Elements.InfoBox.SetText("Search mode. Press <ESC> to exit, <Enter> to go to the search results, C to clear the results")
		t.app.SetFocus(t.search)
	case 'C':
		t.Elements.InfoBox.SetText("Cleared search")
		go t.handleClearSearchCmd()
	case '?':
		t.handleShowHelp()
	}
	return event
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
		t.Elements.Output.SetText("")
		t.Elements.InfoBox.SetText("Running...")
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
		t.Elements.Output.SetText("")
		t.Elements.InfoBox.SetText("Running all tests...")
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
			t.Elements.InfoBox.SetText("No failed tests to run. Good job ")
		})
		return
	}

	var wg sync.WaitGroup

	failedTests := t.state.FailedTests
	t.state.Reset()

	t.app.QueueUpdateDraw(func() {
		t.Elements.Output.SetText("")
		t.Elements.InfoBox.SetText("Running failed tests...")
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
			t.Elements.InfoBox.SetText("No passed tests to run. Try running all tests ")
		})
		return
	}

	var wg sync.WaitGroup

	passedTests := t.state.PassedTests
	t.state.Reset()

	t.app.QueueUpdateDraw(func() {
		t.Elements.Output.SetText("")
		t.Elements.InfoBox.SetText("Running passed tests...")
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
		t.Elements.Output.SetBorderColor(tcell.ColorYellow)
	})

	res := t.runner.Run(test.RunCmd)
	if res.IsSuccess {
		t.app.QueueUpdateDraw(func() {
			t.Elements.Output.SetBorderColor(tcell.ColorGreen)
			testNode.SetText(fmt.Sprintf("[limegreen] [darkturquoise]%s", test.Name))
		})
		t.state.PassedTests = append(t.state.PassedTests, testNode)
	} else {
		t.app.QueueUpdateDraw(func() {
			t.Elements.Output.SetBorderColor(tcell.ColorOrangeRed)
			testNode.SetText(fmt.Sprintf("[orangered] [darkturquoise]%s", test.Name))
		})
		t.state.FailedTests = append(t.state.FailedTests, testNode)
	}

	t.app.QueueUpdateDraw(func() {
		t.Elements.Output.SetText(res.Output)
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
		t.Elements.Output.SetBorderColor(borderColor)
		t.Elements.Output.SetText(output)
	case *models.LazyTest:
		res, ok := t.state.TestOutput[node]
		if ok {
			if res.IsSuccess {
				t.Elements.Output.SetBorderColor(tcell.ColorGreen)
			} else {
				t.Elements.Output.SetBorderColor(tcell.ColorOrangeRed)
			}
			t.Elements.Output.SetText(res.Output)
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

		t.Elements.InfoBox.SetText(msg)
	})
}

func (t *TUI) handleClearSearchCmd() {
	t.app.QueueUpdateDraw(func() {
		t.search.SetText("")
		t.tree.SetRoot(t.state.Root)
		t.app.SetFocus(t.tree)
	})
}
