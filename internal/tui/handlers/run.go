package handlers

import (
	"fmt"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRun(r Runner, a *tview.Application, e *elements.Elements, s *state.State) {
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

	switch v := ref.(type) {
	case *models.LazyTestSuite:
		for _, child := range testNode.GetChildren() {
			test := child.GetReference().(*models.LazyTest)
			go runTest(r, a, e, s, child, test)
		}
		HandleNodeChanged(e, s)(testNode)
	case *models.LazyTest:
		go runTest(r, a, e, s, testNode, v)
	}

	updateRunInfo(a, e, s)
}
