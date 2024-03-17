package handlers

import (
	"fmt"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
)

func (h *Handlers) HandleYankOutput(a tui.Application, c tui.Clipboard, e *elements.Elements) {
	err := c.WriteAll(e.Output.GetText(true))
	a.QueueUpdateDraw(func() {
		if err != nil {
			e.InfoBox.SetText(fmt.Sprintf("Cannot copy output, %v", err))
			return
		}
		e.InfoBox.SetText("Coppied output!")
	})
}
