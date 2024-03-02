package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/loader"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type runner interface {
	Run(command string) *models.LazyTestResult
}

type TUI struct {
	App      *tview.Application
	State    *state.State
	Elements *elements.Elements

	directory string
	runner    runner
	loader    *loader.LazyTestLoader
}

func NewTUI(d string, r runner, e []engines.LazyEngine) *TUI {
	return &TUI{
		App:       tview.NewApplication(),
		State:     state.NewState(),
		directory: d,
		runner:    r,
		loader:    loader.NewLazyTestLoader(e),
	}
}

func (t *TUI) Run() error {
	t.State.TestTree = tview.NewTreeNode(t.directory)
	if err := t.loader.LoadLazyTests(t.directory, t.State.TestTree); err != nil {
		return fmt.Errorf("unable to load tests, %w", err)
	}

	t.Elements = elements.NewElements()
	t.Elements.Setup(
		t.State.TestTree,
		handlers.HandleNodeChanged(t.Elements, t.State),
		handlers.HandleSearchChanged(t.Elements, t.State),
		handlers.HandleSearchDone(t.App, t.Elements, t.State),
		handlers.HandleHelpDone(t.App, t.Elements),
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
	switch pressed_key := event.Rune(); pressed_key {
	case 'q':
		t.App.Stop()
	case '1':
		t.App.SetFocus(t.Elements.Tree)
	case '2':
		t.App.SetFocus(t.Elements.Output)
	case 'r':
		go handlers.HandleRun(t.runner, t.App, t.Elements, t.State)
	case 'a':
		go handlers.HandleRunAll(t.runner, t.App, t.Elements, t.State)
	case 'f':
		go handlers.HandleRunFailed(t.runner, t.App, t.Elements, t.State)
	case 'p':
		go handlers.HandleRunPassed(t.runner, t.App, t.Elements, t.State)
	case '/':
		handlers.HandleSearchFocus(t.App, t.Elements, t.State)
	case 'C':
		go handlers.HandleClearSearch(t.App, t.Elements, t.State)
	case '?':
		t.App.SetRoot(t.Elements.HelpModal, true)
	}
	return event
}
