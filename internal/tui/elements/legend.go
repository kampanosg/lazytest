package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Elements) initLegend() {
	e.Legend.SetBorder(false)
	e.Legend.SetTitleAlign(tview.AlignCenter)
	e.Legend.SetBackgroundColor(tcell.ColorDefault)
	e.Legend.SetText("?: help, 1/2: navigate, /: search, q: quit")
}
