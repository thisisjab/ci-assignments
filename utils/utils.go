package utils

import (
	"math"
	"os"
)

const float64EqualityThreshold = 1e-9

// FloatsEqual makes it easier to compare floats since floats are not really exact.
func FloatsEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
