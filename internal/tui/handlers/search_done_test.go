package handlers_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/mocks"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleSearchDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		App   tui.Application
		Elems *elements.Elements
		State *state.State
	}

	type args struct {
		key tcell.Key
	}

	type want struct {
		searchFieldText string
		infoBoxText     string
		isSearching     bool
	}

	tests := []struct {
		name   string
		fields func() fields
		args   args
		want   want
	}{
		{
			name: "key is not enter or escape",
			fields: func() fields {
				return fields{
					Elems: elements.NewElements(),
					State: state.NewState(),
				}
			},
			args: args{
				key: tcell.KeyBackspace,
			},
			want: want{
				searchFieldText: "",
				infoBoxText:     "",
				isSearching:     false,
			},
		},
		{
			name: "key is enter - focus on results",
			fields: func() fields {
				elems := elements.NewElements()

				mockApp := mocks.NewMockApplication(ctrl)
				mockApp.
					EXPECT().
					SetFocus(elems.Tree).
					Times(1)

				return fields{
					App:   mockApp,
					Elems: elems,
					State: state.NewState(),
				}
			},
			args: args{
				key: tcell.KeyEnter,
			},
			want: want{
				searchFieldText: "",
				infoBoxText:     "",
				isSearching:     false,
			},
		},
		{
			name: "key is escape - leave search",
			fields: func() fields {
				elems := elements.NewElements()

				mockApp := mocks.NewMockApplication(ctrl)
				mockApp.
					EXPECT().
					SetFocus(elems.Tree).
					Times(1)

				return fields{
					App:   mockApp,
					Elems: elems,
					State: state.NewState(),
				}
			},
			args: args{
				key: tcell.KeyEscape,
			},
			want: want{
				searchFieldText: "",
				infoBoxText:     "Exited search mode",
				isSearching:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()

			h := handlers.NewHandlers()
			h.HandleSearchDone(fields.App, fields.Elems, fields.State)(tt.args.key)

			assert.Equal(t, tt.want.searchFieldText, fields.Elems.Search.GetText())
			assert.Equal(t, tt.want.infoBoxText, fields.Elems.InfoBox.GetText(true))
			assert.Equal(t, tt.want.isSearching, fields.State.IsSearching)
		})
	}
}
