package tui_test

import (
	"sync"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/mocks"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"go.uber.org/mock/gomock"
)

func TestInputCapture_HandleRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleRun(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'r', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}

func TestInputCapture_HandleRunAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleRunAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}

func TestInputCapture_HandleRunPassed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleRunPassed(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'p', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}

func TestInputCapture_HandleRunFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleRunFailed(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'f', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}

func TestInputCapture_HandleQuit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mocks.NewMockApplication(ctrl)
	app.EXPECT().Stop().Times(1)

	tui.NewTUI(app, nil, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone))
}

func TestInputCapture_HandleSwitchToTreePane(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := elements.NewElements()
	app := mocks.NewMockApplication(ctrl)
	app.EXPECT().SetFocus(e.Tree).Times(1)

	tui.NewTUI(app, nil, nil, nil, e, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '1', tcell.ModNone))
}

func TestInputCapture_HandleSwitchToOutputPane(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := elements.NewElements()
	app := mocks.NewMockApplication(ctrl)
	app.EXPECT().SetFocus(e.Output).Times(1)

	tui.NewTUI(app, nil, nil, nil, e, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '2', tcell.ModNone))
}

func TestInputCapture_HandleSearchFocus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleSearchFocus(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(1)

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '/', tcell.ModNone))
}

func TestInputCapture_HandleSearchClear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleSearchClear(gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'C', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}

func TestInputCapture_HandleHelp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := elements.NewElements()
	a := mocks.NewMockApplication(ctrl)
	a.EXPECT().SetRoot(e.HelpModal, true).Times(1)

	tui.NewTUI(a, nil, nil, nil, e, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '?', tcell.ModNone))
}

func TestInputCapture_HandleResizeLeft(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleResize(tui.ResizeLeft, gomock.Any(), gomock.Any()).
		Times(1)

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '-', tcell.ModNone))
}

func TestInputCapture_HandleResizeRight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleResize(tui.ResizeRight, gomock.Any(), gomock.Any()).
		Times(1)

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '+', tcell.ModNone))
}

func TestInputCapture_HandleResizeDefault(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleResize(tui.ResizeDefault, gomock.Any(), gomock.Any()).
		Times(1)

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, '0', tcell.ModNone))
}

func TestInputCapture_HandleYankNode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleYankNode(gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'y', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}

func TestInputCapture_HandleYankOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	wg.Add(1)

	h := mocks.NewMockHandlers(ctrl)
	h.EXPECT().
		HandleYankOutput(gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any, _ any) {
			defer wg.Done()
		})

	tui.NewTUI(nil, h, nil, nil, nil, state.NewState(), "", nil).
		InputCapture(tcell.NewEventKey(tcell.KeyRune, 'Y', tcell.ModNone))

	timeout := 5 * time.Second
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("test timeout")
	}
}
