package models

type Weights []float64

type TrainingVectorJsonObject struct {
	Label  float64     `json:"label"`
	Values [][]float64 `json:"values"`
}

type TrainingVector struct {
	T      float64
	Values Weights 
}
