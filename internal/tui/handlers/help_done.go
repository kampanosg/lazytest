package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
)

func HandleHelpDone(a Application, e *elements.Elements) func(btnIdx int, btnLbl string) {
	return func(btnIdx int, btnLbl string) {
		if btnIdx <= 1 {
			a.SetRoot(e.Flex, true)
			a.SetFocus(e.Tree)
		}
	}
}
