package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRunPassed(r runner, a *tview.Application, e *elements.Elements, s *state.State) {
	if len(s.PassedTests) == 0 {
		a.QueueUpdateDraw(func() {
			e.InfoBox.SetText("No passed tests to run. Try running all tests ")
		})
		return
	}

	passedTests := s.PassedTests
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running passed tests...")
	})

	for _, testNode := range passedTests {
		ref := testNode.GetReference()
		if ref == nil {
			continue
		}

		if test, ok := ref.(*models.LazyTest); ok {
			runTest(r, a, e, s, testNode, test)
		}

	}

	updateRunInfo(a, e, s)
}
