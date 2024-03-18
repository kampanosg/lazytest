package handlers_test

import (
	"errors"
	"testing"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/mocks"
	"go.uber.org/mock/gomock"
)

func TestYankOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		App       tui.Application
		Clipboard tui.Clipboard
		Elements  *elements.Elements
	}

	tests := []struct {
		name   string
		fields func() fields
		want   string
	}{
		{
			name: "yank error",
			fields: func() fields {
				c := mocks.NewMockClipboard(ctrl)
				c.EXPECT().
					WriteAll("test").
					Return(errors.New("test error")).
					Times(1)

				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().
					QueueUpdateDraw(gomock.Any()).
					Times(1)

				e := elements.NewElements()
				e.Output.SetText("test")

				return fields{
					App:       a,
					Clipboard: c,
					Elements:  e,
				}
			},
			want: "Cannot copy output, test error",
		},
		{
			name: "success",
			fields: func() fields {
				c := mocks.NewMockClipboard(ctrl)
				c.EXPECT().
					WriteAll("test").
					Return(nil).
					Times(1)

				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().
					QueueUpdateDraw(gomock.Any()).
					Times(1)

				e := elements.NewElements()
				e.Output.SetText("test")

				return fields{
					App:       a,
					Clipboard: c,
					Elements:  e,
				}
			},
			want: "Coppied output!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			h := handlers.NewHandlers()
			h.HandleYankOutput(fields.App, fields.Clipboard, fields.Elements)
		})
	}
}
