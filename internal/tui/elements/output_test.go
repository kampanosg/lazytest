package elements

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	e := NewElements()

	e.Setup(nil, nil, nil, nil, nil)

	assert.Equal(t, "Output", e.Output.GetTitle())
	assert.Equal(t, tcell.ColorDefault, e.Output.GetBackgroundColor())
	assert.Equal(t, "", e.Output.GetText(true))
}
