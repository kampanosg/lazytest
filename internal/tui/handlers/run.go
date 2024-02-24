package handlers

import (
	"fmt"
	"sync"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRun(r runner, a *tview.Application, e *elements.Elements, s *state.State) {
	testNode := e.Tree.GetCurrentNode()
	if testNode == nil {
		return
	}

	ref := testNode.GetReference()
	if ref == nil {
		return
	}

	var wg sync.WaitGroup
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText(fmt.Sprintf("Running %s", testNode.GetText()))
	})

	switch ref.(type) {
	case *models.LazyTestSuite:
		for _, child := range testNode.GetChildren() {
			wg.Add(1)
			test := child.GetReference().(*models.LazyTest)
			runTest(r, a, e, s, &wg, child, test)
		}
		wg.Wait()
		HandleNodeChanged(e, s)(testNode)
	case *models.LazyTest:
		wg.Add(1)
		runTest(r, a, e, s, &wg, testNode, ref.(*models.LazyTest))
		wg.Wait()
	}

	updateRunInfo(a, e, s)
}
