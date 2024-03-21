package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistory(t *testing.T) {
	e := NewElements()

	e.initHistory()

	assert.Equal(t, "History", e.History.GetTitle())
	assert.Equal(t, 0, e.History.GetItemCount())
}