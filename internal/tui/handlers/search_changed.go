package handlers

import (
	"strings"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/internal/tui/utils"
)

func HandleSearchChanged(e *elements.Elements, s *state.State) func(searchQuery string) {
	return func(searchQuery string) {
		if strings.HasSuffix(searchQuery, "/") {
			// when the user presses / to search, the / is still in the input field
			// so we're removing it here
			searchQuery = searchQuery[:len(searchQuery)-1]
			e.Search.SetText(searchQuery)
		}

		if searchQuery == "" {
			e.Tree.SetRoot(s.TestTree)
			return
		}

		root := s.TestTree
		filtered := utils.Search(root, searchQuery)
		e.Tree.SetRoot(filtered)
	}
}
