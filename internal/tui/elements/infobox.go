package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const infoBoxText = "Welcome to LazyTest  "

func (e *Elements) initInfoBox() {
	e.InfoBox.SetBorder(true)
	e.InfoBox.SetTitle("Info")
	e.InfoBox.SetTitleAlign(tview.AlignLeft)
	e.InfoBox.SetBackgroundColor(tcell.ColorDefault)
	e.InfoBox.SetDynamicColors(true)
	e.InfoBox.SetText(infoBoxText)
}
