package elements

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	e := NewElements()

	e.Setup(nil, 4, 8, nil, nil, nil, nil)

	assert.Equal(t, "Search", e.Search.GetTitle())
	assert.Equal(t, tcell.ColorDefault, e.Search.GetBackgroundColor())
	assert.Equal(t, searchPlaceholderStyle, e.Search.GetPlaceholderStyle())
}
