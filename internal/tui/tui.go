package tui

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/loader"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type runner interface {
	Run(command string) *models.LazyTestResult
}

type TUI struct {
	App      *tview.Application
	State    *state.State
	Elements *elements.Elements

	directory string
	runner    runner
	loader    *loader.LazyTestLoader
}

func NewTUI(d string, r runner, e []engines.LazyEngine) *TUI {
	return &TUI{
		App:       tview.NewApplication(),
		State:     state.NewState(),
		directory: d,
		runner:    r,
		loader:    loader.NewLazyTestLoader(e),
	}
}

func (t *TUI) Run() error {
	t.State.TestTree = tview.NewTreeNode(t.directory)
	t.loader.LoadLazyTests(t.directory, t.State.TestTree)

	t.Elements = elements.NewElements(
		t.State.TestTree,
		t.HandleNodeChangedEvent,
		t.handleSearchChangedEvent,
		t.handleSearchDoneEvent,
	)
	t.Elements.Setup()

	t.App.EnableMouse(true)
	t.App.SetInputCapture(t.inputCapture)

	if err := t.App.SetRoot(t.Elements.Flex, true).SetFocus(t.Elements.Tree).EnablePaste(true).Run(); err != nil {
		return err
	}

	return nil
}

func (t *TUI) HandleNodeChangedEvent(node *tview.TreeNode) {
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
			res, ok := t.State.TestOutput[child]
			if ok {
				borderColor = tcell.ColorGreen
				if !res.IsSuccess {
					borderColor = tcell.ColorOrangeRed
				}

				output := fmt.Sprintf("--- %s ---\n%s\n\n", child.GetText(), res.Output)
				outputs = outputs + output
			}
		}
		t.Elements.Output.SetBorderColor(borderColor)
		t.Elements.Output.SetText(outputs)
	case *models.LazyTest:
		res, ok := t.State.TestOutput[node]
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

func (t *TUI) handleSearchChangedEvent(searchQuery string) {
	if strings.HasSuffix(searchQuery, "/") {
		// when the user presses / to search, the / is still in the input field
		// so we're removing it here
		searchQuery = searchQuery[:len(searchQuery)-1]
		t.Elements.Search.SetText(searchQuery)
	}

	if searchQuery == "" {
		t.Elements.Tree.SetRoot(t.State.TestTree)
		return
	}

	root := t.State.TestTree
	filtered := search(root, searchQuery)
	t.Elements.Tree.SetRoot(filtered)
}

func (t *TUI) handleSearchDoneEvent(key tcell.Key) {
	t.State.IsSearching = false

	if key == tcell.KeyEnter {
		t.App.SetFocus(t.Elements.Tree)
	}

	if key == tcell.KeyEscape {
		t.Elements.Search.SetText("")
		t.Elements.Tree.SetRoot(t.State.TestTree)
		t.App.SetFocus(t.Elements.Tree)
		t.Elements.InfoBox.SetText("Exited search mode")
	}
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

func (t *TUI) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	if t.State.IsSearching {
		return event
	}
	switch pressed_key := event.Rune(); pressed_key {
	case 'q':
		t.App.Stop()
	case '1':
		t.App.SetFocus(t.Elements.Tree)
	case '2':
		t.App.SetFocus(t.Elements.Output)
	case 'r':
		go t.handleRunCmd()
	case 'a':
		go t.handleRunAllCmd()
	case 'f':
		go t.handleRunFailedCmd()
	case 'p':
		go t.handleRunPassedCmd()
	case '/':
		t.State.IsSearching = true
		t.Elements.InfoBox.SetText("Search mode. Press <ESC> to exit, <Enter> to go to the search results, C to clear the results")
		t.App.SetFocus(t.Elements.Search)
	case 'C':
		t.Elements.InfoBox.SetText("Cleared search")
		go t.handleClearSearchCmd()
	case '?':
		t.handleShowHelp()
	}
	return event
}

func (t *TUI) handleRunCmd() {
	testNode := t.Elements.Tree.GetCurrentNode()
	if testNode == nil {
		return
	}

	ref := testNode.GetReference()
	if ref == nil {
		return
	}

	var wg sync.WaitGroup
	t.State.Reset()

	t.App.QueueUpdateDraw(func() {
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
	t.State.Reset()

	t.App.QueueUpdateDraw(func() {
		t.Elements.Output.SetText("")
		t.Elements.InfoBox.SetText("Running all tests...")
	})

	t.doRunAll(&wg, t.Elements.Tree.GetRoot().GetChildren())

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
	if len(t.State.FailedTests) == 0 {
		t.App.QueueUpdateDraw(func() {
			t.Elements.InfoBox.SetText("No failed tests to run. Good job ")
		})
		return
	}

	var wg sync.WaitGroup

	failedTests := t.State.FailedTests
	t.State.Reset()

	t.App.QueueUpdateDraw(func() {
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
	if len(t.State.PassedTests) == 0 {
		t.App.QueueUpdateDraw(func() {
			t.Elements.InfoBox.SetText("No passed tests to run. Try running all tests ")
		})
		return
	}

	var wg sync.WaitGroup

	passedTests := t.State.PassedTests
	t.State.Reset()

	t.App.QueueUpdateDraw(func() {
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

	t.App.QueueUpdateDraw(func() {
		testNode.SetText(fmt.Sprintf("[yellow] [darkturquoise]%s", test.Name))
		t.Elements.Output.SetBorderColor(tcell.ColorYellow)
	})

	res := t.runner.Run(test.RunCmd)
	if res.IsSuccess {
		t.App.QueueUpdateDraw(func() {
			t.Elements.Output.SetBorderColor(tcell.ColorGreen)
			testNode.SetText(fmt.Sprintf("[limegreen] [darkturquoise]%s", test.Name))
		})
		t.State.PassedTests = append(t.State.PassedTests, testNode)
	} else {
		t.App.QueueUpdateDraw(func() {
			t.Elements.Output.SetBorderColor(tcell.ColorOrangeRed)
			testNode.SetText(fmt.Sprintf("[orangered] [darkturquoise]%s", test.Name))
		})
		t.State.FailedTests = append(t.State.FailedTests, testNode)
	}

	t.App.QueueUpdateDraw(func() {
		t.Elements.Output.SetText(res.Output)
	})

	t.State.TestOutput[testNode] = res
}

func (t *TUI) handleShowHelp() {
	t.Elements.HelpModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex <= 1 {
			t.App.SetRoot(t.Elements.Flex, true).SetFocus(t.Elements.Tree)
		}
	})
	t.App.SetRoot(t.Elements.HelpModal, true)
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
			res, ok := t.State.TestOutput[child]
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
		res, ok := t.State.TestOutput[node]
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
	t.App.QueueUpdateDraw(func() {
		totalFailed := len(t.State.FailedTests)
		totalPassed := len(t.State.PassedTests)
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
	t.App.QueueUpdateDraw(func() {
		t.Elements.Search.SetText("")
		t.Elements.Tree.SetRoot(t.State.TestTree)
		t.App.SetFocus(t.Elements.Tree)
	})
}
