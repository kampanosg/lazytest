package handlers_test

import (
	"testing"

	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestHandleSearchChanged_EmptyQuery(t *testing.T) {
	e := elements.NewElements()
	s := state.NewState()

	testTree := tview.NewTreeNode("root")

	s.TestTree = testTree

	handlers.HandleSearchChanged(e, s)("")

	assert.Equal(t, testTree, e.Tree.GetRoot())
}

func TestHandleSearchChanged_Query(t *testing.T) {
	e := elements.NewElements()
	s := state.NewState()

	testTree := tview.NewTreeNode("app_test.go")
	testTree.SetReference(&models.LazyTestSuite{
		Path: "TestThisFunction",
	})

	s.TestTree = testTree

	handlers.HandleSearchChanged(e, s)("/Function")

	assert.Equal(t, "Function", e.Search.GetText())
	assert.Len(t, e.Tree.GetRoot().GetChildren(), 1)
}
