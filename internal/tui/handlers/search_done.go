package handlers

import (
	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
)

func HandleSearchDone(a Application, e *elements.Elements, s *state.State) func(key tcell.Key) {
	return func(key tcell.Key) {
		s.IsSearching = false

		if key == tcell.KeyEnter {
			a.SetFocus(e.Tree)
		}

		if key == tcell.KeyEscape {
			e.Search.SetText("")
			e.Tree.SetRoot(s.TestTree)
			e.InfoBox.SetText("Exited search mode")
			a.SetFocus(e.Tree)
		}
	}
}
