package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/rivo/tview"
)

func HandleHelpDone(a *tview.Application, e *elements.Elements) func(btnIdx int, btnLbl string) {
	return func(btnIdx int, btnLbl string) {
		if btnIdx <= 1 {
			a.SetRoot(e.Flex, true).SetFocus(e.Tree)
		}
	}
}
