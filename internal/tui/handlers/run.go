package handlers

import (
	"fmt"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
)

func (h *Handlers) HandleRun(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	testNode := e.Tree.GetCurrentNode()
	if testNode == nil {
		return
	}

	ref := testNode.GetReference()
	if ref == nil {
		return
	}

	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText(fmt.Sprintf("Running %s", testNode.GetText()))
	})

	ch := make(chan *runResult)

	go receiveTestResults(ch, a, e, s, h.HandleNodeChanged)

	switch v := ref.(type) {
	case *models.LazyTestSuite:
		for _, child := range testNode.GetChildren() {
			test := child.GetReference().(*models.LazyTest)
			go runTest(ch, r, a, e, child, test)
		}
	case *models.LazyTest:
		go runTest(ch, r, a, e, testNode, v)
	}
}
