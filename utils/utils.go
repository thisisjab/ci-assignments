package utils

import (
    "fmt"
    "math"
)

func PrintMatrix(input [][]float64) {
	for _, row := range input {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Printf("\n")
	} 
}

const float64EqualityThreshold = 1e-9

// FloatsEqual makes it easier to compare floats since floats are not really exact.
func FloatsEqual(a, b float64) bool {
	return math.Abs(a - b) <= float64EqualityThreshold
}