package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
)

func (h *Handlers) HandleRunFailed(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	if len(s.FailedTests) == 0 {
		a.QueueUpdateDraw(func() {
			e.InfoBox.SetText("No failed tests to run. Good job ")
		})
		return
	}

	failedTests := s.FailedTests
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running failed tests...")
	})

	for _, testNode := range failedTests {
		ref := testNode.GetReference()
		if ref == nil {
			continue
		}

		if test, ok := ref.(*models.LazyTest); ok {
			go runTest(r, a, e, s, testNode, test)
		}
	}

	updateRunInfo(a, e, s)
}
