package handlers

import (
	"fmt"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/pkg/models"
)

func (h *Handlers) HandleYankNode(a tui.Application, c tui.Clipboard, e *elements.Elements) {
	node := e.Tree.GetCurrentNode()
	ref := node.GetReference()

	if ref == nil {
		return
	}

	var err error
	switch v := ref.(type) {
	case *models.LazyTestSuite:
		err = c.WriteAll(v.Path)
	case *models.LazyTest:
		err = c.WriteAll(v.Name)
	default:
		return
	}

	a.QueueUpdateDraw(func() {
		if err != nil {
			e.InfoBox.SetText(fmt.Sprintf("[red] Cannot copy value, %v", err))
			return
		}
		e.InfoBox.SetText("Coppied!")
	})
}
