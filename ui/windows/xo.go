package windows

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"xo-detection/models"
	"xo-detection/training"
	"xo-detection/training/perceptron"
)

var vbox = container.NewVBox()
var values = make(models.Weights, 25)
var buttons [25]*widget.Button
var resultButton *widget.Button
var clearButton *widget.Button
var resultLabel *widget.Label

func init() {
	calculatedWeights, _ := training.GetOrCalculateWeights("xo-perceptron", "saved_weights/xo-perceptron.json")

	// Initiating values
	for i, _ := range values {
		values[i] = -1
	}

	resultLabel = widget.NewLabel("Click on result button to see the result.")
	clearButton = widget.NewButton("Clear", func() { clearValues() })
	resultButton = widget.NewButton("Result", func() {
		calculateResult(calculatedWeights.Weights, calculatedWeights.Bias, calculatedWeights.Theta)
	})
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
	for i, _ := range values {
		values[i] = -1
		buttons[i].SetText("")
		resultLabel.SetText("Click on result button to see the result")
	}
}

func calculateResult(weights models.Weights, bias, theta float64) {
	f := perceptron.Result(values, weights, bias, theta)
	result := ""
	formattedText := "Perceptron: %v"

	if f == 1 {
		result = "X"
	} else {
		result = "O"
	}

	resultLabel.SetText(fmt.Sprintf(formattedText, result))
}
