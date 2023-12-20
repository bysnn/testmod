package util

import "math"

func Swap(a, b int) (int, int) {
	return b, a
}

func Max(a, b int) int {
	big := math.Max(float64(a), float64(b))
	return int(big)
}
