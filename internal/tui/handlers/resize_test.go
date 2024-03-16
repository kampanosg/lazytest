package handlers_test

import (
	"testing"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/state"
)

func TestHandleResize(t *testing.T) {
	type fields struct {
		d tui.ResizeDirection
		e *elements.Elements
		s *state.State
	}

	type want struct {
		sidebarSize     int
		mainContentSize int
	}

	tests := []struct {
		name   string
		fields func() fields
		want   want
	}{
		{
			name: "resize left too small",
			fields: func() fields {
				s := state.NewState()
				s.Size = &state.Size{
					Sidebar:     2,
					MainContent: 10,
				}
				return fields{
					d: tui.ResizeLeft,
					e: elements.NewElements(),
					s: s,
				}
			},
			want: want{
				sidebarSize:     2,
				mainContentSize: 10,
			},
		},
		{
			name: "resize right too small",
			fields: func() fields {
				s := state.NewState()
				s.Size = &state.Size{
					Sidebar:     10,
					MainContent: 2,
				}
				return fields{
					d: tui.ResizeRight,
					e: elements.NewElements(),
					s: s,
				}
			},
			want: want{
				sidebarSize:     10,
				mainContentSize: 2,
			},
		},
		{
			name: "resize right",
			fields: func() fields {
				return fields{
					d: tui.ResizeRight,
					e: elements.NewElements(),
					s: state.NewState(),
				}
			},
			want: want{
				sidebarSize:     5,
				mainContentSize: 7,
			},
		},
		{
			name: "resize left",
			fields: func() fields {
				return fields{
					d: tui.ResizeLeft,
					e: elements.NewElements(),
					s: state.NewState(),
				}
			},
			want: want{
				sidebarSize:     3,
				mainContentSize: 9,
			},
		},
		{
			name: "reset flex",
			fields: func() fields {
				s := state.NewState()
				s.Size = &state.Size{
					Sidebar:     6,
					MainContent: 6,
				}
				return fields{
					d: tui.ResizeDefault,
					e: elements.NewElements(),
					s: s,
				}
			},
			want: want{
				sidebarSize:     4,
				mainContentSize: 8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			h := handlers.NewHandlers()
			h.HandleResize(fields.d, fields.e, fields.s)
		})
	}
}
