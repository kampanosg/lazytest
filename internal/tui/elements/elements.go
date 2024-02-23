package elements

import "github.com/rivo/tview"

type Elements struct {
	InfoBox *tview.TextView
	Legend  *tview.TextView
}

func NewElements() *Elements {
	return &Elements{
		InfoBox: tview.NewTextView(),
		Legend:  tview.NewTextView(),
	}
}

func (e *Elements) Setup() {
	e.initInfoBox()
	e.initLegend()
}
