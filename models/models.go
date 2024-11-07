package models

// Weights vector is a row-matrix: A slice (dynamic array) of floating point numbers.
type Weights []float64

// TrainingVectorJsonObject is just the format of our input JSON file that we expect.
type TrainingVectorJsonObject struct {
	Label  float64   `json:"label"`
	Values []Weights `json:"values"`
}

// TrainingVector consists of an expected value (T) and weights.
type TrainingVector struct {
	T      float64
	Values Weights
}

type SavedWeightAndBiasJsonObject struct {
	Key                  string  `json:"key"`
	Bias                 float64 `json:"bias"`
	Weights              Weights `json:"values"`
	ThetaOrStopCondition float64 `json:"theta_or_stop_condition"`
	LearningRate         float64 `json:"learning_rate"`
	TotalEpoches         int     `json:"total_epoches"`
	TrainingDataSize     int     `json:"training_data_size"`
	TestDataSize         int     `json:"test_data_size"`
	SuccessRate          float64 `json:"success_rate"`
}
