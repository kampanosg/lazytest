package handlers

import (
	"sync"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRunAll(r runner, a *tview.Application, e *elements.Elements, s *state.State) {
	var wg sync.WaitGroup
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running all tests...")
	})

	doRunAll(r, a, e, s, &wg, e.Tree.GetRoot().GetChildren())

	wg.Wait()
	updateRunInfo(a, e, s)
}

func doRunAll(
	r runner,
	a *tview.Application,
	e *elements.Elements,
	s *state.State,
	wg *sync.WaitGroup,
	nodes []*tview.TreeNode,
) {
	for _, testNode := range nodes {
		if len(testNode.GetChildren()) > 0 {
			doRunAll(r, a, e, s, wg, testNode.GetChildren())
		} else {
			ref := testNode.GetReference()
			if ref == nil {
				continue
			}

			switch ref.(type) {
			case *models.LazyTest:
				wg.Add(1)
				runTest(r, a, e, s, wg, testNode, ref.(*models.LazyTest))
			}
		}
	}
}
