package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRunAll(r runner, a *tview.Application, e *elements.Elements, s *state.State) {
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running all tests...")
	})

	doRunAll(r, a, e, s, e.Tree.GetRoot().GetChildren())

	updateRunInfo(a, e, s)
}

func doRunAll(
	r runner,
	a *tview.Application,
	e *elements.Elements,
	s *state.State,
	nodes []*tview.TreeNode,
) {
	for _, testNode := range nodes {
		if len(testNode.GetChildren()) > 0 {
			doRunAll(r, a, e, s, testNode.GetChildren())
		} else {
			ref := testNode.GetReference()
			if ref == nil {
				continue
			}

			switch ref.(type) {
			case *models.LazyTest:
				go runTest(r, a, e, s, testNode, ref.(*models.LazyTest))
			}
		}
	}
}
