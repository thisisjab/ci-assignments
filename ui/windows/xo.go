package windows

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"xo-detection/data"
	"xo-detection/models"
	"xo-detection/training"
	"xo-detection/training/hebb"
	"xo-detection/training/perceptron"
)

var vbox = container.NewVBox()
var values = make(models.Weights, 25)
var buttons [25]*widget.Button
var resultButton *widget.Button
var clearButton *widget.Button
var resultLabel *widget.Label

var perceptronDataObject models.SavedWeightAndBiasJsonObject
var hebbDataObject models.SavedWeightAndBiasJsonObject

func init() {
	// Initiating values
	for i := range values {
		values[i] = -1
	}

	resultLabel = widget.NewLabel("Click on result button to see the result.")
	clearButton = widget.NewButton("Clear", func() { clearValues() })
}

func XOWindow(parent *fyne.Window) *fyne.Container {
	vbox.Add(widget.NewLabel("XO detection with Hebb's rule, perceptron network, and adaline."))
	vbox.Add(widget.NewLabel("I used about 470 vectors to train the various perceptron."))

	// Initializing buttons
	// This also could be done in init function, but would require another loop in XOWindow func. to add buttons to
	// window.
	for i := 0; i < 5; i++ {
		buttonsRow := container.NewGridWithColumns(5)

		for j := 0; j < 5; j++ {
			button := widget.NewButton("", func() {
				toggleValues(i, j)
			})

			buttons[(i*5)+j] = button
			buttonsRow.Add(button)
		}

		vbox.Add(buttonsRow)
	}

	loadWeightsToMemory()

	resultButton = widget.NewButton("Result", func() {
		calculateResult()
	})

	vbox.Add(resultButton)
	vbox.Add(clearButton)
	vbox.Add(resultLabel)

	return vbox
}

func toggleValues(i int, j int) {
	index := (i * 5) + j
	if values[index] == -1 {
		values[index] = 1
		buttons[index].SetText("#")
	} else {
		values[index] = -1
		buttons[index].SetText("")
	}
}

func clearValues() {
	for i := range values {
		values[i] = -1
		buttons[i].SetText("")
		resultLabel.SetText("Click on result button to see the result")
	}
}

func calculateResult() {
	formattedText := "Perceptron: %v | Hebb: %v"

	fPerceptron := perceptron.Result(values, perceptronDataObject.Weights, perceptronDataObject.Bias, perceptronDataObject.Theta)
	resultPerceptron := ""

	fHebb := hebb.Result(values, hebbDataObject.Weights, hebbDataObject.Bias)
	resultHebb := ""

	if fPerceptron == 1 {
		resultPerceptron = "X"
	} else {
		resultPerceptron = "O"
	}

	if fHebb >= 1 {
		resultHebb = "X"
	} else {
		resultHebb = "O"
	}

	resultLabel.SetText(fmt.Sprintf(formattedText, resultPerceptron, resultHebb))
}

// loadWeightsToMemory reads saved weights file or creates the file and saves the weights after training.
func loadWeightsToMemory() {
	perceptronDataObject = training.LoadWeightsOrTrain("saved_weights/xo-perceptron.json", func() models.SavedWeightAndBiasJsonObject {
		loadedTrainingVectors := data.UnmarshalTrainingDataFile("training/data/xo.json")
		cleanedData := data.PrepareData(loadedTrainingVectors)

		initialWeights := make([]models.Weights, 1)
		initialWeights[0] = make(models.Weights, 25)
		initialBias := 0.0
		theta := 0.2
		learningRate := 0.3

		totalEpochs := perceptron.Train(cleanedData, &initialWeights, &initialBias, theta, learningRate)

		return models.SavedWeightAndBiasJsonObject{
			Weights:          initialWeights[len(initialWeights)-1],
			Bias:             initialBias,
			Theta:            theta,
			LearningRate:     learningRate,
			TotalEpoches:     totalEpochs,
			TrainingDataSize: len(cleanedData),
			Key:              "xo-perceptron",
		}
	})

	hebbDataObject = training.LoadWeightsOrTrain("saved_weights/xo-hebb.json", func() models.SavedWeightAndBiasJsonObject {
		loadedTrainingVectors := data.UnmarshalTrainingDataFile("training/data/xo.json")
		cleanedData := data.PrepareData(loadedTrainingVectors)

		initialWeights := make([]models.Weights, 1)
		initialWeights[0] = make(models.Weights, 25)
		initialBias := 0.0

		hebb.Train(cleanedData, &initialWeights, &initialBias)

		return models.SavedWeightAndBiasJsonObject{
			Weights:          initialWeights[len(initialWeights)-1],
			Bias:             initialBias,
			Theta:            -1,
			LearningRate:     -1,
			TotalEpoches:     1,
			TrainingDataSize: len(cleanedData),
			Key:              "xo-perceptron",
		}
	})
}
