package hebb

import (
	m "xo-detection/models"
)

// calculateWeights returns weights and bias
func calculateWeights(v m.TrainingVector, weights m.Weights, oldBias float64) (finalWeights m.Weights, finalBias float64) {
	var calculatedWeights m.Weights

	for i, x := range v.Values {
		delta := x * v.T
		newWeight := delta + weights[i]
		calculatedWeights = append(calculatedWeights, newWeight)
	}

	bias := oldBias + v.T

	return calculatedWeights, bias
}

// Train function gets a initialized weights slices and bias and trains the network keeping track of all weights change.
func Train(vectors []m.TrainingVector, weights *m.Weights, bias *float64) {
	if len(vectors) == 0 {
		panic("Vectors are uninitialized.")
	}

	if len(*weights) == 0 {
		panic("Weights are uninitialized.")
	}

	if len(vectors[0].Values) != len(*weights) {
		panic("Length of weights and values do not match.")
	}

	for _, v := range vectors {
		previousWeights := *weights
		newWeights, newBias := calculateWeights(v, previousWeights, *bias)
		*weights = newWeights
		*bias = newBias
	}
}

// Result function determines the output based on given input using the trained weights slice + final bias.
func Result(inputs, weights m.Weights, bias float64) int8 {
	result := bias

	for i, input := range inputs {
		result += input * weights[i]
	}

	if result > 0 {
		return 1
	}

	return -1
}

func TestSuccessRate(vectors []m.TrainingVector, weights m.Weights, bias float64) float64 {
	successCount := 0

	for _, vector := range vectors {
		result := Result(vector.Values, weights, bias)

		if float64(result) == vector.T {
			successCount++
		}
	}

	successRate := float64(successCount) / float64(len(vectors))

	return successRate
}
