package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func (h *Handlers) HandleRunAll(r tui.Runner, a tui.Application, e *elements.Elements, s *state.State) {
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running all tests...")
	})

	ch := make(chan *runResult)

	go receiveTestResults(ch, a, e, s)

	doRunAll(ch, r, a, e, e.Tree.GetRoot().GetChildren())
}

func doRunAll(
	ch chan<- *runResult,
	r tui.Runner,
	a tui.Application,
	e *elements.Elements,
	nodes []*tview.TreeNode,
) {
	for _, testNode := range nodes {
		if len(testNode.GetChildren()) > 0 {
			doRunAll(ch, r, a, e, testNode.GetChildren())
		} else {
			ref := testNode.GetReference()
			if ref == nil {
				continue
			}
			switch ref := ref.(type) {
			case *models.LazyTest:
				go runTest(ch, r, a, e, testNode, ref)
			}
		}
	}
}
