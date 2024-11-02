package perceptron

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

// CalculateWeights returns weights and bias + a boolean value that determines if weights changed or not.
func CalculateWeights(v m.TrainingVector, weights m.Weights, bias, theta, learningRete float64) (m.Weights, float64, bool) {
	// Net Input = bias + (x1 * w1) + (x2 * w2) + ... + (xn + wn)
	netInput := bias

	for i, x := range v.Values {
		netInput += x * weights[i]
	}

	f := ActivationFunction(netInput, theta)

	if utils.FloatsEqual(v.T, float64(f)) {
		// y(NetInput) is equal to T that we expected.
		return weights, bias, false
	}

	// If code reached here, means weights need to get updated.

	// new weight = old weight + weight delta
	// weight delta = learning rate * x(i) * expected T

	var newWeights m.Weights
	for i, x := range v.Values {
		delta := learningRete * x * v.T
		newWeights = append(newWeights, delta+weights[i])
	}

	// new bias = old bias + bias delta
	// bias delta = learning rate * 1 * expected T
	newBias := bias + (learningRete * 1 * v.T)
	return newWeights, newBias, true
}

// Train function gets a initialized weights slices and bias and trains the network keeping track of all weights change.
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
	currentIteration := 0  // for each epoch
	unmodifiedWeights := 0 // for each epoch

	for { // For non-gophers: This is a while loop.

		for _, v := range vectors {
			lastWeights := (*weights)[len(*weights)-1]
			newWeights, newBias, isChanged := CalculateWeights(v, lastWeights, *bias, theta, learningRate)

			*weights = append(*weights, newWeights)
			*bias = newBias

			if !isChanged {
				unmodifiedWeights += 1
			}

			currentIteration += 1
		}

		if currentIteration >= epochLength {
			// Here, when hit last iteration of current epoch, we check if count of unchanged weights is
			// equal to epoch length. If these two values are equal, it means training must be stopped.

			totalEpochs += 1
			currentIteration = 0

			if unmodifiedWeights >= epochLength {
				return totalEpochs
			}

			unmodifiedWeights = 0
		}
	}
}

// Result function determines the output based on given input using the trained weights slice + final bias.
func Result(inputs, weights m.Weights, bias, theta float64) int8 {
	// Net Input = bias + (x1 * w2) + (x2 * w2) + ... + (xn + wn)
	netInput := bias

	for i, v := range inputs {
		netInput += v * weights[i]
	}

	return ActivationFunction(netInput, theta)
}
