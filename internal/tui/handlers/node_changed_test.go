package handlers_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestHandleNodeChanged(t *testing.T) {
	test1 := &models.LazyTest{
		Name: "LazyTest1",
	}

	test2 := &models.LazyTest{
		Name: "LazyTest2",
	}

	testSuite := &models.LazyTestSuite{
		Path: "Lazy/LazyTest",
		Tests: []*models.LazyTest{
			test1,
			test2,
		},
	}

	testNode1 := tview.NewTreeNode("LazyTest1")
	testNode1.SetReference(test1)

	testNode2 := tview.NewTreeNode("LazyTest2")
	testNode2.SetReference(test2)

	suiteNode := tview.NewTreeNode("Lazy/LazyTest").
		AddChild(testNode1).
		AddChild(testNode2).
		SetReference(testSuite)

	type fields struct {
		Elems *elements.Elements
		State *state.State
	}

	type args struct {
		Node *tview.TreeNode
	}

	type want struct {
		text  string
		title string
		color tcell.Color
	}

	tests := []struct {
		name   string
		fields func() fields
		args   args
		want   want
	}{
		{
			name: "node is a folder",
			fields: func() fields {
				return fields{
					Elems: elements.NewElements(),
					State: state.NewState(),
				}
			},
			args: args{
				Node: tview.NewTreeNode("folder"),
			},
			want: want{
				text:  "",
				title: "",
				color: tcell.ColorWhite,
			},
		},
		{
			name: "node is passed test",
			fields: func() fields {
				s := state.NewState()
				s.TestOutput[testNode1] = &models.LazyTestResult{
					IsSuccess: true,
					Output:    "test passed",
				}

				return fields{
					Elems: elements.NewElements(),
					State: s,
				}
			},
			args: args{
				Node: testNode1,
			},
			want: want{
				text:  "test passed",
				color: tcell.ColorGreen,
				title: "Output - LazyTest1",
			},
		},
		{
			name: "node is failed test",
			fields: func() fields {
				s := state.NewState()
				s.TestOutput[testNode1] = &models.LazyTestResult{
					IsSuccess: false,
					Output:    "test failed",
				}

				return fields{
					Elems: elements.NewElements(),
					State: s,
				}
			},
			args: args{
				Node: testNode1,
			},
			want: want{
				text:  "test failed",
				color: tcell.ColorOrangeRed,
				title: "Output - LazyTest1",
			},
		},
		{
			name: "node is test suite",
			fields: func() fields {
				s := state.NewState()
				s.TestOutput[testNode1] = &models.LazyTestResult{
					IsSuccess: false,
					Output:    "test failed",
				}
				s.TestOutput[testNode2] = &models.LazyTestResult{
					IsSuccess: true,
					Output:    "test passed",
				}
				return fields{
					Elems: elements.NewElements(),
					State: s,
				}
			},
			args: args{
				Node: suiteNode,
			},
			want: want{
				text:  "--- LazyTest1 ---\ntest failed\n\n--- LazyTest2 ---\ntest passed\n\n",
				color: tcell.ColorOrangeRed,
				title: "Output - Lazy/LazyTest",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()

			h := handlers.NewHandlers()
			h.HandleNodeChanged(fields.Elems, fields.State)(tt.args.Node)

			assert.Equal(t, tt.want.text, fields.Elems.Output.GetText(true))
			assert.Equal(t, tt.want.color, fields.Elems.Output.GetBorderColor())
			assert.Equal(t, tt.want.title, fields.Elems.Output.GetTitle())
		})
	}
}
