package handlers

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func (h *Handlers) HandleNodeChanged(e *elements.Elements, s *state.State) func(node *tview.TreeNode) {
	return func(node *tview.TreeNode) {
		node.SetColor(tcell.ColorBlueViolet)

		ref := node.GetReference()
		if ref == nil {
			return
		}

		e.Output.SetBorderColor(tcell.ColorWhite)
		e.Output.SetText("")

		switch v := ref.(type) {
		case *models.LazyTestSuite:
			outputs := ""
			hasTestOutput := false
			hasFailedTest := false
			e.Output.SetBorderColor(tcell.ColorWhite)
			for _, child := range node.GetChildren() {
				res, ok := s.TestOutput[child]
				if ok {
					hasTestOutput = true
					if !res.IsSuccess {
						hasFailedTest = true
					}
					output := fmt.Sprintf("--- %s ---\n%s\n\n", child.GetText(), res.Output)
					outputs = outputs + output
				}
			}

			e.Output.SetText(outputs)
			e.Output.SetTitle(fmt.Sprintf("Output - %s", v.Path))

			e.History.Clear()

			if hasTestOutput {
				if hasFailedTest {
					e.Output.SetBorderColor(tcell.ColorOrangeRed)
				} else {
					e.Output.SetBorderColor(tcell.ColorGreen)
				}
			}
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
			e.Output.SetTitle(fmt.Sprintf("Output - %s", v.Name))
			updateHistory(e, s, node)
			updateTimings(e, s, node)
		}
	}
}

func updateHistory(e *elements.Elements, s *state.State, node *tview.TreeNode) {
	e.History.Clear()
	history := s.History[node]
	idx := 1
	for i := len(history) - 1; i >= 0; i-- {
		item := history[i]
		var txt = fmt.Sprintf("%d.", idx)
		if item.Passed {
			txt = fmt.Sprintf("%s [limegreen]", txt)
		} else {
			txt = fmt.Sprintf("%s [orangered]", txt)
		}

		txt = fmt.Sprintf("%s %v", txt, item.Timestamp.Format("2006-01-02 @ 15:04:05"))
		e.History.AddItem(txt, "", 0, nil)
		idx += 1
	}
}

func updateTimings(e *elements.Elements, s *state.State, node *tview.TreeNode) {
	e.Timings.Clear()
	timings := s.Timings[node]
	idx := 1
	for i := len(timings) - 1; i >= 0; i-- {
		color := "[white]"
		item := timings[i]

		var next time.Duration
		if i-1 >= 0 {
			next = timings[i-1]
		}

		if next == 0 {
		} else if next.Milliseconds() >= item.Milliseconds() {
			color = "[limegreen]"
		} else {
			color = "[orangered]"
		}

		txt := fmt.Sprintf("%s%d. %vms", color, idx, item.Milliseconds())
		e.Timings.AddItem(txt, "", 0, nil)
		idx += 1
	}
}

