package handlers_test

import (
	"testing"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/handlers/mocks"
	"go.uber.org/mock/gomock"
)

func TestHandleHelpDone(t *testing.T) {
	ctrl := gomock.NewController(t)

	type fields struct {
		App   handlers.Application
		Elems *elements.Elements
	}

	type args struct {
		BtnIndex int
		BtnLbl   string
	}

	tests := []struct {
		name   string
		fields func() fields
		args   args
	}{
		{
			name: "close help",
			fields: func() fields {
				mockApp := mocks.NewMockApplication(ctrl)
				mockApp.
					EXPECT().
					SetRoot(gomock.Any(), true).
					Times(1)
				mockApp.
					EXPECT().
					SetFocus(gomock.Any()).
					Times(1)

				return fields{
					App:   mockApp,
					Elems: elements.NewElements(),
				}
			},
			args: args{
				BtnIndex: -1,
				BtnLbl:   "",
			},
		},
		{
			name: "dont close help",
			fields: func() fields {
				mockApp := mocks.NewMockApplication(ctrl)
				mockApp.
					EXPECT().
					SetRoot(gomock.Any(), true).
					Times(0)
				mockApp.
					EXPECT().
					SetFocus(gomock.Any()).
					Times(0)

				return fields{
					App:   mockApp,
					Elems: elements.NewElements(),
				}
			},
			args: args{
				BtnIndex: 6,
				BtnLbl:   "Enter",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			handlers.HandleHelpDone(fields.App, fields.Elems)(tt.args.BtnIndex, tt.args.BtnLbl)
		})
	}
}
