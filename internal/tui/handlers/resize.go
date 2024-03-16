package handlers

import (
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/state"
)

func (h *Handlers) HandleResize(d tui.ResizeDirection, e *elements.Elements, s *state.State) {
	switch d {
	case tui.ResizeLeft:
		handleResizeLeft(e, s)
	case tui.ResizeRight:
		handleResizeRight(e, s)
	case tui.ResizeDefault:
		handleResizeDefault(e, s)
	}
}

func handleResizeLeft(e *elements.Elements, s *state.State) {
	if s.Size.Sidebar == 2 {
		return
	}

	s.Size.Sidebar -= 1
	s.Size.MainContent += 1

	e.ResizeFlex(s.Size.Sidebar, s.Size.MainContent)
}

func handleResizeRight(e *elements.Elements, s *state.State) {
	if s.Size.MainContent == 2 {
		return
	}

	s.Size.Sidebar += 1
	s.Size.MainContent -= 1

	e.ResizeFlex(s.Size.Sidebar, s.Size.MainContent)
}

func handleResizeDefault(e *elements.Elements, s *state.State) {
	s.Size.Sidebar = 4
	s.Size.MainContent = 8
	e.ResizeFlex(s.Size.Sidebar, s.Size.MainContent)
}
