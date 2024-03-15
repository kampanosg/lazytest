package elements

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestInfoBox(t *testing.T) {
	e := NewElements()

	e.Setup(nil, 4, 8, nil, nil, nil, nil)

	assert.Equal(t, "Info", e.InfoBox.GetTitle())
	assert.Equal(t, tcell.ColorDefault, e.InfoBox.GetBackgroundColor())
	assert.Equal(t, infoBoxText, e.InfoBox.GetText(true))
}
