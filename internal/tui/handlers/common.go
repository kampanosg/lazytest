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

func updateRunInfo(a tui.Application, e *elements.Elements, s *state.State) {
	a.QueueUpdateDraw(func() {
		totalFailed := len(s.FailedTests)
		totalPassed := len(s.PassedTests)
		msg := "Finished running."
		if totalPassed > 0 {
			msg = fmt.Sprintf("%s [limegreen]%d passed", msg, totalPassed)
		}
		if totalFailed > 0 {
			msg = fmt.Sprintf("%s. [orangered]%d failed", msg, totalFailed)
		}

		e.InfoBox.SetText(msg)
	})
}

func runTest(
	r tui.Runner,
	a tui.Application,
	e *elements.Elements,
	s *state.State,
	testNode *tview.TreeNode,
	test *models.LazyTest,
) {
	a.QueueUpdateDraw(func() {
		testNode.SetText(fmt.Sprintf("[yellow] [darkturquoise]%s", test.Name))
		e.Output.SetBorderColor(tcell.ColorYellow)
	})

	res := r.Run(test.RunCmd)
	if res.IsSuccess {
		a.QueueUpdateDraw(func() {
			e.Output.SetBorderColor(tcell.ColorGreen)
			testNode.SetText(fmt.Sprintf("[limegreen] [darkturquoise]%s", test.Name))
		})
		s.PassedTests = append(s.PassedTests, testNode)
	} else {
		a.QueueUpdateDraw(func() {
			e.Output.SetBorderColor(tcell.ColorOrangeRed)
			testNode.SetText(fmt.Sprintf("[orangered] [darkturquoise]%s", test.Name))
		})
		s.FailedTests = append(s.FailedTests, testNode)
	}

	a.QueueUpdateDraw(func() {
		e.Output.SetText(res.Output)
	})

	s.TestOutput[testNode] = res
}
