package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/rivo/tview"
)

//go:generate mockgen -source=$GOFILE -destination=mocks/$GOFILE -package=mocks
type Application interface {
	SetRoot(root tview.Primitive, fullscreen bool) *tview.Application
	SetFocus(p tview.Primitive) *tview.Application
}

func HandleHelpDone(a Application, e *elements.Elements) func(btnIdx int, btnLbl string) {
	return func(btnIdx int, btnLbl string) {
		if btnIdx <= 1 {
			a.SetRoot(e.Flex, true)
			a.SetFocus(e.Tree)
		}
	}
}
