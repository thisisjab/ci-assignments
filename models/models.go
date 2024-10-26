package models

// Each weights vector is a row-matrix: A slice (dynamic array) of floating point numbers.
type Weights []float64

// This is just the format of our input JSON file that we expect.
type TrainingVectorJsonObject struct {
	Label  float64     `json:"label"`
	Values [][]float64 `json:"values"`
}

// Each training vector consists of an expected value (T) and weights.
type TrainingVector struct {
	T      float64
	Values Weights 
}
