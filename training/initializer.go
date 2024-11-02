package training

import (
	"log"
	"xo-detection/data"
	"xo-detection/models"
	"xo-detection/training/perceptron"
	"xo-detection/utils"
)

func initPerceptronWeights(path string) (calculatedWeights models.SavedWeightAndBiasJsonObject, calculated bool) {
	if utils.FileExists(path) {
		log.Println("Loading perceptron weights from " + path)
		loadedData := data.UnmarshalWeightsFile(path)
		log.Printf(
			"Tooks %v epoches for %v to get trained with %v data and theta=%v and learning rate=%v.\n",
			loadedData.TotalEpoches, "xo-perceptron", loadedData.TrainingDataSize, loadedData.Theta, loadedData.LearningRate)
		return loadedData, false
	}

	log.Println("File for perceptron weights does not exist. Training the network...")
	log.Println("Initializing initial weights with 0.")
	var weights []models.Weights
	var firstRowWeights models.Weights
	for i := 0; i < 25; i++ {
		firstRowWeights = append(firstRowWeights, 0)
	}

	weights = append(weights, firstRowWeights)

	var bias float64

	const theta = 0.2
	const learningRate = 0.3

	log.Printf("Using %v as theta and %v as learning rate.\n", theta, learningRate)

	loadedTrainingData := data.UnmarshalTrainingDataFile("./training/data/xo.json")
	log.Printf("Data for xo is read from %v.\n", path)

	cleanedData := data.PrepareData(loadedTrainingData)
	log.Printf("Data for xo is cleaned from %v.\n", path)
	log.Printf("Loaded %v vectors as training data.\n", len(loadedTrainingData))

	totalEpochs := perceptron.Train(cleanedData, &weights, &bias, theta, learningRate)

	result := models.SavedWeightAndBiasJsonObject{Weights: weights[len(weights)-1],
		Key:              "xo-perceptron",
		TotalEpoches:     totalEpochs,
		Bias:             bias,
		Theta:            theta,
		LearningRate:     learningRate,
		TrainingDataSize: len(cleanedData),
	}

	log.Printf(
		"Tooks %v epoches for %v to get trained with %v data and theta=%v and learning rate=%v.\n",
		result.TotalEpoches, "xo-perceptron", result.TrainingDataSize, result.Theta, result.LearningRate)

	success, err := data.SaveWeights(result, path)

	if !success {
		log.Fatalln(err)
	}

	return result, true
}

func GetOrCalculateWeights(key, path string) (weightsAndBias models.SavedWeightAndBiasJsonObject, calculated bool) {
	switch key {
	case "xo-perceptron":
		return initPerceptronWeights(path)
	default:
		panic("Unhandled key for `GetOrCalculateWeights`")
	}
}
