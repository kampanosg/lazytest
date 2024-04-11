package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimings(t *testing.T) {
	e := NewElements()

	e.initTimings()

	assert.Equal(t, "Timings", e.Timings.GetTitle())
	assert.Equal(t, 0, e.Timings.GetItemCount())
}
