package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Elements) initOutput() {
	e.Output.SetBorder(true)
	e.Output.SetTitle("Output")
	e.Output.SetTitleAlign(tview.AlignLeft)
	e.Output.SetBackgroundColor(tcell.ColorDefault)
	e.Output.SetScrollable(true)
	e.Output.SetDynamicColors(true)
	e.Output.SetRegions(true)
}
