package handlers

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleNodeChanged(e *elements.Elements, s *state.State) func(node *tview.TreeNode) {
	return func(node *tview.TreeNode) {
		node.SetColor(tcell.ColorBlueViolet)

		ref := node.GetReference()
		if ref == nil {
			return
		}

		switch ref.(type) {
		case *models.LazyTestSuite:
			borderColor := tcell.ColorWhite
			outputs := ""
			for _, child := range node.GetChildren() {
				res, ok := s.TestOutput[child]
				if ok {
					borderColor = tcell.ColorGreen
					if !res.IsSuccess {
						borderColor = tcell.ColorOrangeRed
					}

					output := fmt.Sprintf("--- %s ---\n%s\n\n", child.GetText(), res.Output)
					outputs = outputs + output
				}
			}
			e.Output.SetBorderColor(borderColor)
			e.Output.SetText(outputs)
		case *models.LazyTest:
			res, ok := s.TestOutput[node]
			if ok {
				if res.IsSuccess {
					e.Output.SetBorderColor(tcell.ColorGreen)
				} else {
					e.Output.SetBorderColor(tcell.ColorOrangeRed)
				}
				e.Output.SetText(res.Output)
			}
		}
	}
}
