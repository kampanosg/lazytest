package handlers_test

import (
	"testing"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/mocks"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleSearchFocus(t *testing.T) {
	ctrl := gomock.NewController(t)

	type fields struct {
		App   tui.Application
		Elems *elements.Elements
		State *state.State
	}

	type want struct {
		isSearching bool
		infoBoxText string
	}

	tests := []struct {
		name   string
		fields func() fields
		want   want
	}{
		{
			name: "search mode enabled",
			fields: func() fields {
				elems := elements.NewElements()
				mockApp := mocks.NewMockApplication(ctrl)
				mockApp.
					EXPECT().
					SetFocus(elems.Search).
					Times(1)

				return fields{
					App:   mockApp,
					Elems: elems,
					State: state.NewState(),
				}
			},
			want: want{
				isSearching: true,
				infoBoxText: "Search mode. Press <ESC> to exit, <Enter> to go to the search results, C to clear the results",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			
			h:= handlers.NewHandlers()
			h.HandleSearchFocus(fields.App, fields.Elems, fields.State)

			assert.Equal(t, tt.want.isSearching, fields.State.IsSearching)
			assert.Equal(t, tt.want.infoBoxText, fields.Elems.InfoBox.GetText(true))
		})
	}
}
