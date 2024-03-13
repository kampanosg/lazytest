package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
)

func (h *Handlers) HandleSearchClear(a tui.Application, e *elements.Elements, s *state.State) {
	a.QueueUpdateDraw(func() {
		e.Search.SetText("")
		e.Tree.SetRoot(s.TestTree)
		e.InfoBox.SetText("Cleared search")
		a.SetFocus(e.Tree)
	})
}
