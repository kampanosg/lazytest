package elements

import "github.com/rivo/tview"

type Elements struct {
	InfoBox *tview.TextView
}

func NewElements() *Elements {
	return &Elements{
		InfoBox: tview.NewTextView(),
	}
}

func (e *Elements) Setup() {
	e.initInfoBox()
}
