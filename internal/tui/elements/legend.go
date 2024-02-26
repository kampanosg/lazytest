package elements

import (
	"github.com/gdamore/tcell/v2"
)

const legendText = "?: help, 1/2: navigate, /: search, q: quit"

func (e *Elements) initLegend() {
	e.Legend.SetBorder(false)
	e.Legend.SetBackgroundColor(tcell.ColorDefault)
	e.Legend.SetText(legendText)
}
