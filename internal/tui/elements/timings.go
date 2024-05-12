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
	e.Timings.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'k':
				e.Timings.SetCurrentItem(e.Timings.GetCurrentItem() - 1)
			case 'j':
				currentItem := e.Timings.GetCurrentItem() + 1
				if e.Timings.GetCurrentItem() == e.Timings.GetItemCount()-1 {
					currentItem = 0
				}
				e.Timings.SetCurrentItem(currentItem)
			}
		}
		return event
	})
}
