package handlers_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestHandleNodeChanged(t *testing.T) {
	e := elements.NewElements()
	s := state.NewState()
	node := tview.NewTreeNode("folder")

	handlers.HandleNodeChanged(e, s)(node)

	assert.Equal(t, e.Output.GetText(true), "")
	assert.Equal(t, tcell.ColorWhite, e.Output.GetBorderColor())
}
