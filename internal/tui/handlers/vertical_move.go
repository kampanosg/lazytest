package handlers

import "github.com/kampanosg/lazytest/internal/tui/elements"

func (h *Handlers) HandleMoveUp(e *elements.Elements) {
	if e.History.HasFocus() {
		e.History.SetCurrentItem(e.History.GetCurrentItem() - 1)
		return
	}

	if e.Timings.HasFocus() {
		e.Timings.SetCurrentItem(e.Timings.GetCurrentItem() - 1)
		return
	}
}

func (h *Handlers) HandleMoveDown(e *elements.Elements) {
	// When at the last index, tview won't move the selection to the top when moving down further
	// This is the exact oposite behaviour when moving up
	// So we have to reset the index to zero
	if e.History.HasFocus() {
		currentItem := e.History.GetCurrentItem() + 1
		if currentItem == e.History.GetItemCount()-1 {
			currentItem = 0
		}
		e.History.SetCurrentItem(currentItem)
	}

	if e.Timings.HasFocus() {
		currentItem := e.Timings.GetCurrentItem() + 1
		if currentItem == e.Timings.GetItemCount()-1 {
			currentItem = 0
		}
		e.Timings.SetCurrentItem(currentItem)
	}
}
