package windows

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"xo-detection/data"
	"xo-detection/models"
	"xo-detection/training"
	"xo-detection/training/adaline"
	"xo-detection/training/hebb"
	"xo-detection/training/perceptron"
)

var vbox = container.NewVBox()
var values = make(models.Weights, 25)
var buttons [25]*widget.Button
var resultButton *widget.Button
var clearButton *widget.Button
var resultLabel *widget.Label

var hebbSuccessRateLabel *widget.Label
var perceptronSuccessRateLabel *widget.Label
var perceptronLearningRateEntry *widget.Entry
var perceptronThetaEntry *widget.Entry
var adalineSuccessRateLabel *widget.Label
var adalineLearningRateEntry *widget.Entry
var adalineStopConditionEntry *widget.Entry

const CountOfDataToUse = 400

var perceptronDataObject models.SavedWeightAndBiasJsonObject
var adalineDataObject models.SavedWeightAndBiasJsonObject
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

	hebbSuccessRateLabel = widget.NewLabel("Hebb Success Rate: ")

	perceptronSuccessRateLabel = widget.NewLabel("Perceptron Success Rate: ")
	perceptronLearningRateEntry = widget.NewEntry()
	perceptronLearningRateEntry.SetText("0.1")
	perceptronThetaEntry = widget.NewEntry()
	perceptronThetaEntry.SetText("0.5")

	adalineSuccessRateLabel = widget.NewLabel("Adaline Success Rate: ")
	adalineStopConditionEntry = widget.NewEntry()
	adalineStopConditionEntry.SetText("0.0001")
	adalineLearningRateEntry = widget.NewEntry()
	adalineLearningRateEntry.SetText("0.0001")

	loadWeightsToMemory(true, true)

	resultButton = widget.NewButton("Result", func() {
		calculateResult()
	})

	vbox.Add(resultButton)
	vbox.Add(clearButton)
	vbox.Add(resultLabel)

	vbox.Add(widget.NewSeparator())

	trainingLabelsGrid := container.NewGridWithColumns(4)
	trainingLabelsGrid.Add(widget.NewLabel(""))
	trainingLabelsGrid.Add(widget.NewLabel("Learning Rate"))
	trainingLabelsGrid.Add(widget.NewLabel("Theta / Stop Condition"))
	trainingLabelsGrid.Add(widget.NewLabel("Success Rate"))
	vbox.Add(trainingLabelsGrid)

	vbox.Add(widget.NewSeparator())

	hebbGrid := container.NewGridWithColumns(4)
	hebbGrid.Add(widget.NewLabel("Hebb"))
	hebbGrid.Add(widget.NewLabel(""))
	hebbGrid.Add(widget.NewLabel(""))
	hebbGrid.Add(hebbSuccessRateLabel)
	vbox.Add(hebbGrid)

	vbox.Add(widget.NewSeparator())

	// Entries for training perceptron
	perceptronGrid := container.NewGridWithColumns(4)
	perceptronGrid.Add(widget.NewLabel("Perceptron"))
	perceptronLearningRateEntry.SetPlaceHolder("Learning Rate")
	perceptronThetaEntry.SetPlaceHolder("Theta")
	perceptronGrid.Add(perceptronLearningRateEntry)
	perceptronGrid.Add(perceptronThetaEntry)
	perceptronGrid.Add(perceptronSuccessRateLabel)
	vbox.Add(perceptronGrid)

	vbox.Add(widget.NewSeparator())

	// Entries for training adaline
	adalineGrid := container.NewGridWithColumns(4)
	adalineGrid.Add(widget.NewLabel("Adaline"))
	adalineLearningRateEntry.SetPlaceHolder("Learning Rate")
	adalineStopConditionEntry.SetPlaceHolder("Stop Condition")
	adalineGrid.Add(adalineLearningRateEntry)
	adalineGrid.Add(adalineStopConditionEntry)
	adalineGrid.Add(adalineSuccessRateLabel)
	vbox.Add(adalineGrid)

	vbox.Add(widget.NewSeparator())

	trainGrid := container.NewGridWithColumns(3)
	trainGrid.Add(widget.NewLabel("You can train the network here."))
	trainGrid.Add(widget.NewLabel(""))
	trainGrid.Add(widget.NewButton("Start Training", func() {
		go loadWeightsToMemory(false, false)
	}))
	vbox.Add(trainGrid)

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
	formattedText := "Perceptron: %v | Adaline: %v | Hebb: %v"

	fPerceptron := perceptron.Result(values, perceptronDataObject.Weights, perceptronDataObject.Bias, perceptronDataObject.ThetaOrStopCondition)
	resultPerceptron := ""

	fAdaline := adaline.Result(values, adalineDataObject.Weights, adalineDataObject.Bias)
	resultAdaline := ""

	fHebb := hebb.Result(values, hebbDataObject.Weights, hebbDataObject.Bias)
	resultHebb := ""

	if fPerceptron == 1 {
		resultPerceptron = "X"
	} else {
		resultPerceptron = "O"
	}

	if fAdaline == 1 {
		resultAdaline = "X"
	} else {
		resultAdaline = "O"
	}

	if fHebb >= 1 {
		resultHebb = "X"
	} else {
		resultHebb = "O"
	}

	resultLabel.SetText(fmt.Sprintf(formattedText, resultPerceptron, resultAdaline, resultHebb))
}

// loadWeightsToMemory reads saved weights file or creates the file and saves the weights after training.
func loadWeightsToMemory(saveToDisk, loadFromDisk bool) {

	perceptronThetaEntry.Disable()
	perceptronLearningRateEntry.Disable()
	perceptronSuccessRateLabel.SetText("Calculating")
	adalineStopConditionEntry.Disable()
	adalineLearningRateEntry.Disable()
	adalineSuccessRateLabel.SetText("Calculating")

	perceptronDataObject = training.LoadWeightsOrTrain("saved_weights/xo-perceptron.json", saveToDisk, loadFromDisk, func() models.SavedWeightAndBiasJsonObject {
		loadedTrainingVectors := data.UnmarshalTrainingDataFile("training/data/xo_shuffled_data.json")
		cleanedData := data.PrepareData(loadedTrainingVectors)
		trainingData := cleanedData[:CountOfDataToUse]
		testData := cleanedData[CountOfDataToUse:]

		initialWeights := make(models.Weights, 25)
		initialBias := 0.0

		theta, _ := strconv.ParseFloat(perceptronThetaEntry.Text, 64)
		learningRate, _ := strconv.ParseFloat(perceptronLearningRateEntry.Text, 64)

		totalEpochs := perceptron.Train(trainingData, &initialWeights, &initialBias, theta, learningRate)

		return models.SavedWeightAndBiasJsonObject{
			Weights:              initialWeights,
			Bias:                 initialBias,
			ThetaOrStopCondition: theta,
			LearningRate:         learningRate,
			TotalEpoches:         totalEpochs,
			TrainingDataSize:     len(trainingData),
			TestDataSize:         len(testData),
			SuccessRate:          perceptron.TestSuccessRate(testData, initialWeights, initialBias, theta),
			Key:                  "xo-perceptron",
		}
	})

	adalineDataObject = training.LoadWeightsOrTrain("saved_weights/xo-adaline.json", saveToDisk, loadFromDisk, func() models.SavedWeightAndBiasJsonObject {
		loadedTrainingVectors := data.UnmarshalTrainingDataFile("training/data/xo_shuffled_data.json")
		cleanedData := data.PrepareData(loadedTrainingVectors)
		trainingData := cleanedData[:CountOfDataToUse]
		testData := cleanedData[CountOfDataToUse:]

		initialWeights := make(models.Weights, 25)
		initialBias := 0.0
		stopCondition, _ := strconv.ParseFloat(adalineStopConditionEntry.Text, 64)
		learningRate, _ := strconv.ParseFloat(adalineLearningRateEntry.Text, 64)

		totalEpochs := adaline.Train(trainingData, &initialWeights, &initialBias, learningRate, stopCondition)

		return models.SavedWeightAndBiasJsonObject{
			Weights:              initialWeights,
			Bias:                 initialBias,
			ThetaOrStopCondition: stopCondition,
			LearningRate:         learningRate,
			TotalEpoches:         totalEpochs,
			TrainingDataSize:     len(trainingData),
			TestDataSize:         len(testData),
			SuccessRate:          adaline.TestSuccessRate(testData, initialWeights, initialBias),
			Key:                  "xo-adaline",
		}
	})

	hebbDataObject = training.LoadWeightsOrTrain("saved_weights/xo-hebb.json", saveToDisk, loadFromDisk, func() models.SavedWeightAndBiasJsonObject {
		loadedTrainingVectors := data.UnmarshalTrainingDataFile("training/data/xo_shuffled_data.json")
		cleanedData := data.PrepareData(loadedTrainingVectors)
		trainingData := cleanedData[:CountOfDataToUse]
		testData := cleanedData[CountOfDataToUse:]

		initialWeights := make(models.Weights, 25)
		initialBias := 0.0

		hebb.Train(trainingData, &initialWeights, &initialBias)

		return models.SavedWeightAndBiasJsonObject{
			Weights:              initialWeights,
			Bias:                 initialBias,
			ThetaOrStopCondition: -1,
			LearningRate:         -1,
			TotalEpoches:         1,
			TrainingDataSize:     len(trainingData),
			TestDataSize:         len(testData),
			SuccessRate:          hebb.TestSuccessRate(testData, initialWeights, initialBias),
			Key:                  "xo-hebb",
		}
	})

	perceptronThetaEntry.Enable()
	perceptronLearningRateEntry.Enable()
	adalineStopConditionEntry.Enable()
	adalineLearningRateEntry.Enable()

	hebbSuccessRateLabel.SetText(strconv.FormatFloat(hebbDataObject.SuccessRate, 'f', 6, 64))
	perceptronThetaEntry.SetText(strconv.FormatFloat(perceptronDataObject.ThetaOrStopCondition, 'f', 6, 64))
	perceptronLearningRateEntry.SetText(strconv.FormatFloat(perceptronDataObject.LearningRate, 'f', 6, 64))
	perceptronSuccessRateLabel.SetText(strconv.FormatFloat(perceptronDataObject.SuccessRate, 'f', 6, 64))

	adalineStopConditionEntry.SetText(strconv.FormatFloat(adalineDataObject.ThetaOrStopCondition, 'f', 6, 64))
	adalineLearningRateEntry.SetText(strconv.FormatFloat(adalineDataObject.LearningRate, 'f', 6, 64))
	adalineSuccessRateLabel.SetText(strconv.FormatFloat(adalineDataObject.SuccessRate, 'f', 6, 64))
}
