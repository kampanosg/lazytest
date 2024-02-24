package elements

import (
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/rivo/tview"
)

type Elements struct {
	State     *state.State
	Tree      *tview.TreeView
	Output    *tview.TextView
	Search    *tview.InputField
	InfoBox   *tview.TextView
	Legend    *tview.TextView
	HelpModal *tview.Modal
}

func NewElements(s *state.State) *Elements {
	return &Elements{
		State:     s,
		Tree:      tview.NewTreeView(),
		Output:    tview.NewTextView(),
		InfoBox:   tview.NewTextView(),
		Search:    tview.NewInputField(),
		Legend:    tview.NewTextView(),
		HelpModal: tview.NewModal(),
	}
}

func (e *Elements) Setup() {
	e.initTree()
	e.initOutput()
	e.initInfoBox()
	e.initSearch()
	e.initLegend()
	e.initHelp()
}
