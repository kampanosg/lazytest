package handlers_test

import (
	"testing"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/handlers/mocks"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleSearchFocus(t *testing.T) {
	ctrl := gomock.NewController(t)

	type fields struct {
		App   handlers.Application
		Elems *elements.Elements
		State *state.State
	}

	tests := []struct {
		name          string
		fields        func() fields
		wantSearching bool
		wantText      string
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
			wantSearching: true,
			wantText:      "Search mode. Press <ESC> to exit, <Enter> to go to the search results, C to clear the results",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			handlers.HandleSearchFocus(fields.App, fields.Elems, fields.State)
			assert.Equal(t, tt.wantSearching, fields.State.IsSearching)
			assert.Equal(t, tt.wantText, fields.Elems.InfoBox.GetText(true))
		})
	}
}
