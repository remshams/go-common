package styles

import "math"

func CalculateDimensionsFromPercentage(percent int, total int) int {
	return int(math.Round(float64(total) * (float64(percent) / 100)))
}
