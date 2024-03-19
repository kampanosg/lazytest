package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
)

func (h *Handlers) HandleRunPassed(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
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

	ch := make(chan *runResult)

	go receiveTestResults(ch, a, e, s, h.HandleNodeChanged)

	for _, testNode := range passedTests {
		ref := testNode.GetReference()
		if ref == nil {
			continue
		}

		if test, ok := ref.(*models.LazyTest); ok {
			go runTest(ch, r, a, e, testNode, test)
		}
	}
}
