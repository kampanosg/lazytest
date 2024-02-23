package elements

import "github.com/rivo/tview"

type Elements struct {
	Output    *tview.TextView
	InfoBox   *tview.TextView
	Legend    *tview.TextView
	HelpModal *tview.Modal
}

func NewElements() *Elements {
	return &Elements{
		Output:    tview.NewTextView(),
		InfoBox:   tview.NewTextView(),
		Legend:    tview.NewTextView(),
		HelpModal: tview.NewModal(),
	}
}

func (e *Elements) Setup() {
	e.initOutput()
	e.initInfoBox()
	e.initLegend()
	e.initHelp()
}
