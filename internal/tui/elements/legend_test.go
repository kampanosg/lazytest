package elements

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestLegend(t *testing.T) {
	e := NewElements()

	e.Setup(nil, 4, 8, nil, nil, nil, nil)

	assert.Equal(t, legendText, e.Legend.GetText(true))
	assert.Equal(t, tcell.ColorDefault, e.Legend.GetBackgroundColor())
}
