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

func TestHandleSearchClear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		App   tui.Application
		Elems *elements.Elements
		State *state.State
	}

	type want struct {
		searchFieldText string
		infoBoxText     string
	}

	tests := []struct {
		name   string
		fields func() fields
		want   want
	}{
		{
			name: "search results cleared",
			fields: func() fields {
				mockApp := mocks.NewMockApplication(ctrl)
				mockApp.
					EXPECT().
					QueueUpdateDraw(gomock.Any())

				return fields{
					App:   mockApp,
					Elems: elements.NewElements(),
					State: state.NewState(),
				}
			},
			want: want{
				searchFieldText: "",
				infoBoxText:     "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			h := handlers.NewHandlers()
			h.HandleSearchClear(fields.App, fields.Elems, fields.State)
			assert.Equal(t, tt.want.searchFieldText, fields.Elems.Search.GetText())
			assert.Equal(t, tt.want.infoBoxText, fields.Elems.InfoBox.GetText(true))
		})
	}
}
