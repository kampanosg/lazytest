package handlers

import (
	"fmt"

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
		testNode.SetText(fmt.Sprintf("[yellow] [darkturquoise]%s", test.Name))
		e.Output.SetBorderColor(tcell.ColorYellow)
	})

	res := r.Run(test.RunCmd)
	ch <- &runResult{
		node: testNode,
		res:  res,
		test: test,
	}
}

func receiveTestResults(ch <-chan *runResult, a tui.Application, e *elements.Elements, s *state.State) {
	for {
		res := <-ch
		handleTestFinished(a, e, s, res)
	}
}

func handleTestFinished(a tui.Application, e *elements.Elements, s *state.State, testResult *runResult) {
	txt := fmt.Sprintf("[orangered] [darkturquoise]%s", testResult.test.Name)
	borderColor := tcell.ColorOrangeRed
	if testResult.res.IsSuccess {
		txt = fmt.Sprintf("[limegreen] [darkturquoise]%s", testResult.test.Name)
		borderColor = tcell.ColorGreen
	}

	s.TestOutput[testResult.node] = testResult.res

	if testResult.res.IsSuccess {
		s.PassedTests = append(s.PassedTests, testResult.node)
	} else {
		s.FailedTests = append(s.FailedTests, testResult.node)
	}

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
	})

}
