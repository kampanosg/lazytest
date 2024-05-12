package handlers_test

import (
	"testing"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestHandleMoveUp_History(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.History.AddItem("test 1", "", 0, nil)
	elems.History.AddItem("test 2", "", 0, nil)
	elems.History.AddItem("test 3", "", 0, nil)
	elems.History.SetCurrentItem(2)

	elems.Timings.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.History)

	h.HandleMoveUp(elems)

	assert.Equal(t, 1, elems.History.GetCurrentItem())
	assert.Equal(t, 0, elems.Timings.GetCurrentItem())
}

func TestHandleMoveUp_HistoryInTopElement(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.History.AddItem("test 1", "", 0, nil)
	elems.History.AddItem("test 2", "", 0, nil)
	elems.History.SetCurrentItem(0)

	elems.Timings.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.History)

	h.HandleMoveUp(elems)

	assert.Equal(t, 1, elems.History.GetCurrentItem())
	assert.Equal(t, 0, elems.Timings.GetCurrentItem())
}

func TestHandleMoveUp_Timings(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.Timings.AddItem("test 1", "", 0, nil)
	elems.Timings.AddItem("test 2", "", 0, nil)
	elems.Timings.AddItem("test 3", "", 0, nil)
	elems.Timings.SetCurrentItem(2)

	elems.History.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.Timings)

	h.HandleMoveUp(elems)

	assert.Equal(t, 1, elems.Timings.GetCurrentItem())
	assert.Equal(t, 0, elems.History.GetCurrentItem())
}

func TestHandleMoveUp_TimingsInTopElement(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.Timings.AddItem("test 1", "", 0, nil)
	elems.Timings.AddItem("test 2", "", 0, nil)
	elems.Timings.SetCurrentItem(0)

	elems.History.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.Timings)

	h.HandleMoveUp(elems)

	assert.Equal(t, 1, elems.Timings.GetCurrentItem())
	assert.Equal(t, 0, elems.History.GetCurrentItem())
}

func TestHandleMoveDown_History(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.History.AddItem("test 1", "", 0, nil)
	elems.History.AddItem("test 2", "", 0, nil)
	elems.History.AddItem("test 3", "", 0, nil)
	elems.History.SetCurrentItem(1)

	elems.Timings.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.History)

	h.HandleMoveDown(elems)

	assert.Equal(t, 2, elems.History.GetCurrentItem())
	assert.Equal(t, 0, elems.Timings.GetCurrentItem())
}

func TestHandleMovDown_HistoryInBottomElement(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.History.AddItem("test 1", "", 0, nil)
	elems.History.AddItem("test 2", "", 0, nil)
	elems.History.SetCurrentItem(1)

	elems.Timings.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.History)

	h.HandleMoveDown(elems)

	assert.Equal(t, 0, elems.History.GetCurrentItem())
	assert.Equal(t, 0, elems.Timings.GetCurrentItem())
}

func TestHandleMoveDown_Timings(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.Timings.AddItem("test 1", "", 0, nil)
	elems.Timings.AddItem("test 2", "", 0, nil)
	elems.Timings.AddItem("test 3", "", 0, nil)
	elems.Timings.SetCurrentItem(1)

	elems.History.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.Timings)

	h.HandleMoveDown(elems)

	assert.Equal(t, 2, elems.Timings.GetCurrentItem())
	assert.Equal(t, 0, elems.History.GetCurrentItem())
}

func TestHandleMoveDown_TimingsInBottomElement(t *testing.T) {
	elems := elements.NewElements()
	h := handlers.NewHandlers()
	a := tview.NewApplication()

	elems.Timings.AddItem("test 1", "", 0, nil)
	elems.Timings.AddItem("test 2", "", 0, nil)
	elems.Timings.SetCurrentItem(1)

	elems.History.AddItem("test 4", "", 0, nil)

	a.SetFocus(elems.Timings)

	h.HandleMoveDown(elems)

	assert.Equal(t, 0, elems.Timings.GetCurrentItem())
	assert.Equal(t, 0, elems.History.GetCurrentItem())
}
