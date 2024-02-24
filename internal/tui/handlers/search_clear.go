package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/rivo/tview"
)

func HandleClearSearch(a *tview.Application, e *elements.Elements, s *state.State) {
	a.QueueUpdateDraw(func() {
		e.Search.SetText("")
		e.Tree.SetRoot(s.TestTree)
		e.InfoBox.SetText("Cleared search")
		a.SetFocus(e.Tree)
	})
}
