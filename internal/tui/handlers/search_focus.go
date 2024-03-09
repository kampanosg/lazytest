package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
)

func HandleSearchFocus(a Application, e *elements.Elements, s *state.State) {
	s.IsSearching = true
	e.InfoBox.SetText("Search mode. Press <ESC> to exit, <Enter> to go to the search results, C to clear the results")
	a.SetFocus(e.Search)
}