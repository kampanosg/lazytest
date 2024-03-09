package handlers

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

func HandleRunAll(r Runner, a *tview.Application, e *elements.Elements, s *state.State) {
	s.Reset()

	a.QueueUpdateDraw(func() {
		e.Output.SetText("")
		e.InfoBox.SetText("Running all tests...")
	})

	doRunAll(r, a, e, s, e.Tree.GetRoot().GetChildren())

	updateRunInfo(a, e, s)
}

func doRunAll(
	r Runner,
	a *tview.Application,
	e *elements.Elements,
	s *state.State,
	nodes []*tview.TreeNode,
) {
	for _, testNode := range nodes {
		if len(testNode.GetChildren()) > 0 {
			doRunAll(r, a, e, s, testNode.GetChildren())
		} else {
			ref := testNode.GetReference()
			if ref == nil {
				continue
			}

			switch ref := ref.(type) {
			case *models.LazyTest:
				go func() {
					a.QueueUpdateDraw(func() {
						testNode.SetText(fmt.Sprintf("[yellow] [darkturquoise]%s", ref.Name))
						e.Output.SetBorderColor(tcell.ColorYellow)
					})
					res := r.Run(ref.RunCmd)
					if res.IsSuccess {
						a.QueueUpdateDraw(func() {
							e.Output.SetBorderColor(tcell.ColorGreen)
							testNode.SetText(fmt.Sprintf("[limegreen] [darkturquoise]%s", ref.Name))
						})
						s.PassedTests = append(s.PassedTests, testNode)
					} else {
						a.QueueUpdateDraw(func() {
							e.Output.SetBorderColor(tcell.ColorOrangeRed)
							testNode.SetText(fmt.Sprintf("[orangered] [darkturquoise]%s", ref.Name))
						})
						s.FailedTests = append(s.FailedTests, testNode)
					}
					a.QueueUpdateDraw(func() {
						e.Output.SetText(res.Output)
					})
					s.TestOutput[testNode] = res
				}()
			}
		}
	}
}
