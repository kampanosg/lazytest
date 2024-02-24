package handlers

import (
	"sync"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRunFailed(r runner, a *tview.Application, e *elements.Elements, s *state.State) {
	if len(s.FailedTests) == 0 {
		a.QueueUpdateDraw(func() {
			e.InfoBox.SetText("No failed tests to run. Good job ")
		})
		return
	}

	var wg sync.WaitGroup

	failedTests := s.FailedTests
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running failed tests...")
	})

	for _, testNode := range failedTests {
		wg.Add(1)
		ref := testNode.GetReference()
		if ref == nil {
			continue
		}

		if test, ok := ref.(*models.LazyTest); ok {
			runTest(r, a, e, s, &wg, testNode, test)
		}
	}

	wg.Wait()
	updateRunInfo(a, e, s)
}
