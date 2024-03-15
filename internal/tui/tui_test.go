package tui_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/mocks"
	"github.com/rivo/tview"
	"go.uber.org/mock/gomock"
)

func TestInputCapture(t *testing.T) {
	ctrl := gomock.NewController(t)

	type fields struct {
		App      tui.Application
		Elements *elements.Elements
		Handlers tui.Handlers
	}

	type args struct {
		key *tcell.EventKey
	}

	tests := []struct {
		name   string
		fields func() fields
		args   args
	}{
		{
			name: "handle quit",
			fields: func() fields {
				app := mocks.NewMockApplication(ctrl)
				app.EXPECT().Stop().Times(1)
				return fields{
					App:      app,
					Handlers: handlers.NewHandlers(),
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('q'), tcell.ModNone),
			},
		},
		{
			name: "handle switch to test pane",
			fields: func() fields {
				e := elements.NewElements()
				app := mocks.NewMockApplication(ctrl)
				app.EXPECT().SetFocus(e.Tree).Times(1)
				return fields{
					App:      app,
					Elements: e,
					Handlers: handlers.NewHandlers(),
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('1'), tcell.ModNone),
			},
		},
		{
			name: "handle switch to output pane",
			fields: func() fields {
				e := elements.NewElements()
				app := mocks.NewMockApplication(ctrl)
				app.EXPECT().SetFocus(e.Output).Times(1)
				return fields{
					App:      app,
					Elements: e,
					Handlers: handlers.NewHandlers(),
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('2'), tcell.ModNone),
			},
		},
		{
			name: "handle run",
			fields: func() fields {
				h := mocks.NewMockHandlers(ctrl)
				h.EXPECT().
					HandleRun(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
				return fields{
					App:      tview.NewApplication(),
					Elements: elements.NewElements(),
					Handlers: h,
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyRune, rune('r'), tcell.ModNone),
			},
		},
		{
			name: "handle run all",
			fields: func() fields {
				h := mocks.NewMockHandlers(ctrl)
				h.EXPECT().
					HandleRunAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
				return fields{
					App:      tview.NewApplication(),
					Elements: elements.NewElements(),
					Handlers: h,
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('a'), tcell.ModNone),
			},
		},
		{
			name: "handle run failed",
			fields: func() fields {
				h := mocks.NewMockHandlers(ctrl)
				h.EXPECT().
					HandleRunFailed(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
				return fields{
					App:      tview.NewApplication(),
					Elements: elements.NewElements(),
					Handlers: h,
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('f'), tcell.ModNone),
			},
		},
		{
			name: "handle run passed",
			fields: func() fields {
				h := mocks.NewMockHandlers(ctrl)
				h.EXPECT().
					HandleRunPassed(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
				return fields{
					App:      tview.NewApplication(),
					Elements: elements.NewElements(),
					Handlers: h,
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('p'), tcell.ModNone),
			},
		},
		{
			name: "handle search",
			fields: func() fields {
				h := mocks.NewMockHandlers(ctrl)
				h.EXPECT().
					HandleSearchFocus(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
				return fields{
					App:      tview.NewApplication(),
					Elements: elements.NewElements(),
					Handlers: h,
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('/'), tcell.ModNone),
			},
		},
		{
			name: "handle search clear",
			fields: func() fields {
				h := mocks.NewMockHandlers(ctrl)
				h.EXPECT().
					HandleSearchClear(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
				return fields{
					App:      tview.NewApplication(),
					Elements: elements.NewElements(),
					Handlers: h,
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyCenter, rune('C'), tcell.ModNone),
			},
		},
		{
			name: "handle help",
			fields: func() fields {
				e := elements.NewElements()
				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().SetRoot(e.HelpModal, true).Times(1)
				return fields{
					App:      a,
					Elements: e,
					Handlers: handlers.NewHandlers(),
				}
			},
			args: args{
				key: tcell.NewEventKey(tcell.KeyEnter, rune('?'), tcell.ModNone),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			tui.NewTUI(f.App, f.Handlers, nil, f.Elements, "", nil).
				InputCapture(tt.args.key)
		})
	}

}
