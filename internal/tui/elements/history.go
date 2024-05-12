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
	e.History.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'k':
				e.History.SetCurrentItem(e.History.GetCurrentItem() - 1)
			case 'j':
				currentItem := e.History.GetCurrentItem() + 1
				if e.History.GetCurrentItem() == e.History.GetItemCount()-1 {
					currentItem = 0
				}
				e.History.SetCurrentItem(currentItem)
			}
		}
		return event
	})
}
