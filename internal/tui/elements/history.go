package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Elements) initHistory() {
	e.History.SetBorder(true)
	e.History.SetTitle("History")
	e.History.SetTitleAlign(tview.AlignLeft)
	e.History.SetBackgroundColor(tcell.ColorDefault)
	e.History.ShowSecondaryText(false)
	e.History.SetSelectedBackgroundColor(tcell.ColorBlueViolet)
}
