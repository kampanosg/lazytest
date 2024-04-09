package handlers

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type runResult struct {
	node *tview.TreeNode
	test *models.LazyTest
	res  *models.LazyTestResult
}

func runTest(
	ch chan<- *runResult,
	r tui.Runner,
	a tui.Application,
	e *elements.Elements,
	testNode *tview.TreeNode,
	test *models.LazyTest,
) {
	a.QueueUpdateDraw(func() {
		testNode.SetText(fmt.Sprintf("[yellow]󰦖 [darkturquoise]%s", test.Name))
		e.Output.SetBorderColor(tcell.ColorYellow)
	})

	res := r.RunTest(test.RunCmd)
	ch <- &runResult{
		node: testNode,
		res:  res,
		test: test,
	}
}

func receiveTestResults(ch <-chan *runResult, a tui.Application, e *elements.Elements, s *state.State, hnc func(e *elements.Elements, s *state.State) func(node *tview.TreeNode)) {
	for {
		res := <-ch
		handleTestFinished(a, e, s, res, hnc)
	}
}

func handleTestFinished(a tui.Application, e *elements.Elements, s *state.State, testResult *runResult, hnc func(e *elements.Elements, s *state.State) func(node *tview.TreeNode)) {
	txt := fmt.Sprintf("[orangered] [darkturquoise]%s", testResult.test.Name)
	borderColor := tcell.ColorOrangeRed
	if testResult.res.Passed {
		txt = fmt.Sprintf("[limegreen] [darkturquoise]%s", testResult.test.Name)
		borderColor = tcell.ColorGreen
	}

	s.TestOutput[testResult.node] = testResult.res

	if testResult.res.Passed {
		s.PassedTests = append(s.PassedTests, testResult.node)
	} else {
		s.FailedTests = append(s.FailedTests, testResult.node)
	}

	s.History[testResult.node] = append(s.History[testResult.node], state.HistoricalEntry{
		Timestamp: time.Now(),
		Passed:    testResult.res.Passed,
	})

	s.Timings[testResult.node] = append(s.Timings[testResult.node], testResult.res.Duration)

	a.QueueUpdateDraw(func() {
		testResult.node.SetText(txt)
		e.Output.SetBorderColor(borderColor)
		e.Output.SetText(testResult.res.Output)

		totalPassed := len(s.PassedTests)
		totalFailed := len(s.FailedTests)
		msg := "Finished running"

		if totalPassed > 0 {
			msg = fmt.Sprintf("%s. [limegreen]%d passed", msg, totalPassed)
		}

		if totalFailed > 0 {
			msg = fmt.Sprintf("%s. [orangered]%d failed", msg, totalFailed)
		}

		e.InfoBox.SetText(msg)

		hnc(e, s)(testResult.node)
	})

}
