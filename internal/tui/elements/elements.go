package elements

import "github.com/rivo/tview"

type Elements struct {
	Output  *tview.TextView
	InfoBox *tview.TextView
	Legend  *tview.TextView
}

func NewElements() *Elements {
	return &Elements{
		Output:  tview.NewTextView(),
		InfoBox: tview.NewTextView(),
		Legend:  tview.NewTextView(),
	}
}

func (e *Elements) Setup() {
	e.initOutput()
	e.initInfoBox()
	e.initLegend()
}
