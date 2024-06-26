package tui

//go:generate mockgen -source=$GOFILE -destination=mocks/mocks.go -package=mocks

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/loader"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type Application interface {
	SetRoot(root tview.Primitive, fullscreen bool) *tview.Application
	SetFocus(p tview.Primitive) *tview.Application
	EnableMouse(enable bool) *tview.Application
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Application
	QueueUpdateDraw(f func()) *tview.Application
	Stop()
}

type Runner interface {
	RunTest(command string) *models.LazyTestResult
}

type Handlers interface {
	HandleNodeChanged(e *elements.Elements, s *state.State) func(node *tview.TreeNode)
	HandleHelpDone(a Application, e *elements.Elements) func(btnIdx int, btnLbl string)
	HandleRun(r Runner, a Application, e *elements.Elements, s *state.State)
	HandleRunAll(r Runner, a Application, e *elements.Elements, s *state.State)
	HandleRunFailed(r Runner, a Application, e *elements.Elements, s *state.State)
	HandleRunPassed(r Runner, a Application, e *elements.Elements, s *state.State)
	HandleSearchChanged(e *elements.Elements, s *state.State) func(searchQuery string)
	HandleSearchDone(a Application, e *elements.Elements, s *state.State) func(key tcell.Key)
	HandleSearchFocus(a Application, e *elements.Elements, s *state.State)
	HandleSearchClear(a Application, e *elements.Elements, s *state.State)
	HandleResize(d ResizeDirection, e *elements.Elements, s *state.State)
	HandleYankNode(a Application, c Clipboard, e *elements.Elements)
	HandleYankOutput(a Application, c Clipboard, e *elements.Elements)
}

type Clipboard interface {
	WriteAll(text string) error
}

type ResizeDirection int

const (
	ResizeLeft ResizeDirection = iota
	ResizeRight
	ResizeDefault
)

type TUI struct {
	App       Application
	State     *state.State
	Elements  *elements.Elements
	Handlers  Handlers
	Runner    Runner
	Clipboard Clipboard

	directory string
	loader    *loader.LazyTestLoader
}

func NewTUI(
	a Application,
	h Handlers,
	r Runner,
	c Clipboard,
	e *elements.Elements,
	s *state.State,
	d string,
	eng []engines.LazyEngine,
) *TUI {
	return &TUI{
		App:       a,
		Handlers:  h,
		Runner:    r,
		Clipboard: c,
		State:     s,
		Elements:  e,
		directory: d,
		loader:    loader.NewLazyTestLoader(eng),
	}
}

func (t *TUI) Run() error {
	testNode, err := t.loader.LoadLazyTests(t.directory)
	if err != nil {
		return fmt.Errorf("unable to load tests, %w", err)
	}

	t.State.TestTree = testNode

	t.Elements = elements.NewElements()
	t.Elements.Setup(
		t.State.TestTree,
		t.State.Size.Sidebar,
		t.State.Size.MainContent,
		t.Handlers.HandleNodeChanged(t.Elements, t.State),
		t.Handlers.HandleSearchChanged(t.Elements, t.State),
		t.Handlers.HandleSearchDone(t.App, t.Elements, t.State),
		t.Handlers.HandleHelpDone(t.App, t.Elements),
	)

	t.App.EnableMouse(true)
	t.App.SetInputCapture(t.InputCapture)

	if err := t.App.SetRoot(t.Elements.Flex, true).SetFocus(t.Elements.Tree).EnablePaste(true).Run(); err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}

	return nil
}

func (t *TUI) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	if t.State.IsSearching {
		return event
	}

	switch key := event.Key(); key {
	case tcell.KeyRune:
		switch event.Rune() {
		case 'q':
			t.App.Stop()
		case '1':
			t.App.SetFocus(t.Elements.Tree)
		case '2':
			t.App.SetFocus(t.Elements.Output)
		case '3':
			t.App.SetFocus(t.Elements.History)
		case '4':
			t.App.SetFocus(t.Elements.Timings)
		case 'r':
			go t.Handlers.HandleRun(t.Runner, t.App, t.Elements, t.State)
		case 'a':
			go t.Handlers.HandleRunAll(t.Runner, t.App, t.Elements, t.State)
		case 'f':
			go t.Handlers.HandleRunFailed(t.Runner, t.App, t.Elements, t.State)
		case 'p':
			go t.Handlers.HandleRunPassed(t.Runner, t.App, t.Elements, t.State)
		case '/':
			t.Handlers.HandleSearchFocus(t.App, t.Elements, t.State)
		case 'C':
			go t.Handlers.HandleSearchClear(t.App, t.Elements, t.State)
		case '+':
			t.Handlers.HandleResize(ResizeRight, t.Elements, t.State)
		case '-':
			t.Handlers.HandleResize(ResizeLeft, t.Elements, t.State)
		case '0':
			t.Handlers.HandleResize(ResizeDefault, t.Elements, t.State)
		case 'y':
			go t.Handlers.HandleYankNode(t.App, t.Clipboard, t.Elements)
		case 'Y':
			go t.Handlers.HandleYankOutput(t.App, t.Clipboard, t.Elements)
		case '?':
			t.App.SetRoot(t.Elements.HelpModal, true)
		}
	}
	return event
}
