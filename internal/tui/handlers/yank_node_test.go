package handlers

import (
	"errors"
	"testing"

	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/mocks"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
	"go.uber.org/mock/gomock"
)

func TestHandleYankNode(t *testing.T) {
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
			name: "no ref",
			fields: func() fields {
				node := tview.NewTreeNode("test")
				e := elements.NewElements()
				e.Tree.SetCurrentNode(node)
				return fields{
					App:       nil,
					Clipboard: nil,
					Elements:  e,
				}
			},
			want: "",
		},
		{
			name: "yank test node",
			fields: func() fields {
				test := &models.LazyTest{
					Name: "test",
				}

				node := tview.NewTreeNode("test")
				node.SetReference(test)

				e := elements.NewElements()
				e.Tree.SetCurrentNode(node)

				c := mocks.NewMockClipboard(ctrl)
				c.EXPECT().
					WriteAll("test").
					Return(nil).
					Times(1)

				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().
					QueueUpdateDraw(gomock.Any()).
					Times(1)

				return fields{
					App:       a,
					Clipboard: c,
					Elements:  e,
				}
			},
			want: "Coppied!",
		},
		{
			name: "error when yanking test node",
			fields: func() fields {
				test := &models.LazyTest{
					Name: "test",
				}

				node := tview.NewTreeNode("test")
				node.SetReference(test)

				e := elements.NewElements()
				e.Tree.SetCurrentNode(node)

				c := mocks.NewMockClipboard(ctrl)
				c.EXPECT().
					WriteAll("test").
					Return(errors.New("fubar")).
					Times(1)

				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().
					QueueUpdateDraw(gomock.Any()).
					Times(1)

				return fields{
					App:       a,
					Clipboard: c,
					Elements:  e,
				}
			},
			want: "[red] Cannot copy value, foobar",
		},
		{
			name: "yank test suite",
			fields: func() fields {
				suite := &models.LazyTestSuite{
					Path: "LazyTest/test",
				}

				node := tview.NewTreeNode("test")
				node.SetReference(suite)

				e := elements.NewElements()
				e.Tree.SetCurrentNode(node)

				c := mocks.NewMockClipboard(ctrl)
				c.EXPECT().
					WriteAll("LazyTest/test").
					Return(nil).
					Times(1)

				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().
					QueueUpdateDraw(gomock.Any()).
					Times(1)

				return fields{
					App:       a,
					Clipboard: c,
					Elements:  e,
				}
			},
			want: "Coppied!",
		},
		{
			name: "error when yanking test suite node",
			fields: func() fields {
				suite := &models.LazyTestSuite{
					Path: "LazyTest/test",
				}

				node := tview.NewTreeNode("test")
				node.SetReference(suite)

				e := elements.NewElements()
				e.Tree.SetCurrentNode(node)

				c := mocks.NewMockClipboard(ctrl)
				c.EXPECT().
					WriteAll("LazyTest/test").
					Return(errors.New("fubar")).
					Times(1)

				a := mocks.NewMockApplication(ctrl)
				a.EXPECT().
					QueueUpdateDraw(gomock.Any()).
					Times(1)

				return fields{
					App:       a,
					Clipboard: c,
					Elements:  e,
				}
			},
			want: "[red] Cannot copy value, foobar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			h := NewHandlers()
			h.HandleYankNode(fields.App, fields.Clipboard, fields.Elements)
		})
	}
}
