package adaline

import (
	"math"
	"xo-detection/models"
)

// Train - as the name suggests - does the training for adaline network.
func Train(trainingVector []models.TrainingVector, weights *models.Weights, bias *float64, learningRate, stopCondition float64) int {
	if len(trainingVector) == 0 {
		panic("empty training vector")
	}

	if len(trainingVector[0].Values) != len(*weights) {
		panic("size of weights != size of a single training vector")
	}

	// Used for tracking errors for each epoch
	errorSumList := make([]float64, 0)

	totalEpochs := 0

	for {
		sumOfErrorsForThisEpoch := 0.0

		for _, vector := range trainingVector {
			netInput := *bias

			for i, x := range vector.Values {
				netInput += (*weights)[i] * x
			}

			e := vector.T - netInput

			// Updating weights
			for i, x := range vector.Values {
				(*weights)[i] += learningRate * e * x
			}

			// Updating bias
			*bias += learningRate * e

			sumOfErrorsForThisEpoch += math.Pow(e, 2)
		}

		totalEpochs += 1
		errorSumList = append(errorSumList, sumOfErrorsForThisEpoch)

		// Check the errors only if there are at least two errors
		if len(errorSumList) > 2 {
			differenceOfLastTwoErrors := errorSumList[len(errorSumList)-1] - errorSumList[len(errorSumList)-2]

			if math.Abs(differenceOfLastTwoErrors) <= stopCondition {
				return totalEpochs
			}
		}
	}
}

func Result(inputs, weights models.Weights, bias float64) int8 {
	yNetInput := bias

	for i, x := range inputs {
		yNetInput += x * weights[i]
	}

	if yNetInput > 0 {
		return 1
	}

	return -1
}
