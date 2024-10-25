package utils

import (
    "fmt"
    "math"
)

const float64EqualityThreshold = 1e-9

func PrintMatrix(input [][]float64) {
	for _, row := range input {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Printf("\n")
	} 
}

func FloatsEqual(a, b float64) bool {
	return math.Abs(a - b) <= float64EqualityThreshold
}