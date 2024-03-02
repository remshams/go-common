package styles

import "math"

func CalculateDimensionsFromPercentage(percent int, total int, max *int) int {
	dimensions := int(math.Round(float64(total) * (float64(percent) / 100)))
	if max != nil && dimensions > *max {
		dimensions = *max
	}
	return dimensions
}
