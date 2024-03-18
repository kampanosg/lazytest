package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Elements) initTimings() {
	e.Timings.SetBorder(true)
	e.Timings.SetTitle("Timings")
	e.Timings.SetTitleAlign(tview.AlignLeft)
	e.Timings.SetBackgroundColor(tcell.ColorDefault)
	e.Timings.ShowSecondaryText(false)
	e.Timings.SetSelectedBackgroundColor(tcell.ColorBlueViolet)
}
