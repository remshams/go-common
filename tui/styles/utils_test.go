package styles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDimensionsFromPercentage_ShouldReturnPercentage(t *testing.T) {
	percent := 50
	total := 100
	absolute := 50
	dimensions := CalculateDimensionsFromPercentage(percent, total, UnlimitedDimension)
	assert.Equal(t, absolute, dimensions)
}

func TestCalculateDimensionsFromPercentage_ShouldReturnMax(t *testing.T) {
	percent := 50
	total := 100
	max := 25
	dimensions := CalculateDimensionsFromPercentage(percent, total, max)
	assert.Equal(t, max, dimensions)
}
