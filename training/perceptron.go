package training

import (
	m "xo-detection/models"
	"xo-detection/utils"
)

func ActivationFunction(netInput, theta float64) int8 {
	if theta < netInput {
		return 1
	} else if -theta <= netInput && netInput <= theta {
		return 0
	} else {
		return -1
	}
}

func CalculateWeights(v m.TrainingVector, weights m.Weights, bias, theta, learningRete float64) (m.Weights, float64, bool) {
	netInput := bias

	for i, x := range v.Values {
		netInput += (x * weights[i])
	}

	f := ActivationFunction(netInput, theta)

	if utils.FloatsEqual(v.T, float64(f)) {
		return weights, bias, false
	}

	var newWeights m.Weights
	for i, x := range v.Values {
		delta := learningRete * x * v.T
		newWeights = append(newWeights, (delta + weights[i]))
	}

	newBias := bias + (learningRete * 1 * v.T)
	return newWeights, newBias, true
}

func Train(vectors []m.TrainingVector, weights *[]m.Weights, bias *float64, theta, learningRate float64) int {
	if len(vectors) == 0 {
		panic("Vectors are uninitialized.")
	}

	if len(*weights) == 0 {
		panic("Weights are uninitialized.")
	}

	if len(vectors[0].Values) != len((*weights)[0]) {
		panic("Length of weights and values do not match.")
	}

	epochLength := len(vectors)
	totalEpochs := 0
	currentIteration := 0
	unmodifiedWeights := 0

	for {

		for _, v := range vectors {
			lastWeights := (*weights)[len(*weights)-1]
			newWeights, newBias, isChanged := CalculateWeights(v, lastWeights, *bias, theta, learningRate)

			(*weights) = append((*weights), newWeights)
			*bias = newBias

			if !isChanged {
				unmodifiedWeights += 1
			}

			currentIteration += 1
		}

		if currentIteration >= epochLength {
			totalEpochs += 1
			currentIteration = 0

			if unmodifiedWeights >= epochLength {
				return totalEpochs
			}

			unmodifiedWeights = 0
		}
	}
}

func Result(inputs, weights m.Weights, bias, theta float64) int8 {
	netInput := bias
	for i, v := range inputs {
		netInput += v * weights[i]
	}

	return ActivationFunction(netInput, theta)
}
