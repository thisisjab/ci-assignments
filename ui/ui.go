package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	m "xo-detection/models"
	t "xo-detection/training"
)

func ShowResult(values [25]int, weights m.Weights, bias, theta float64, label *widget.Label) {
	var valuesAsWeights m.Weights

	for _, v := range values {
		valuesAsWeights = append(valuesAsWeights, float64(v))
	} 

	result := t.Result(valuesAsWeights, weights, bias, theta)

	if result == 1 {
		label.SetText( "Result is X")	
	} else {
		label.SetText( "Result is O")
	}
}

func Toggle(x, y int, values *[25]int, buttons [25]*widget.Button) bool {
	index := (x * 5) + y

	if (*values)[index] == -1 {
		(*values)[index] = 1
		buttons[index].SetText("*")
		return true
	}
	(*values)[index] = -1
	buttons[index].SetText("")
	return false
}

func Clear(values *[25]int, buttons [25]*widget.Button, resultLabel *widget.Label) {
	for i := 0; i < 25; i++ {
		values[i] = -1
		buttons[i].SetText("")
	}
	resultLabel.SetText("Draw something and click result button.")
}

func RunUI(weights m.Weights, bias, theta float64) {
	var values [25]int
	var buttons [25]*widget.Button
	for i := range values {
		values[i] = -1
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("XO-Detector")
	mainContainer := container.NewVBox()
	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(500, 500))

	for row := 0; row < 5; row++ {
		colsLayout := container.NewGridWithColumns(5)

		for col := 0; col < 5; col++ {
			check := widget.NewButton("", func() {
				Toggle(row, col, &values, buttons)
			})
			check.Resize(fyne.NewSize(200, 80))
			colsLayout.Add(check)

			buttons[row*5+col] = check
		}

		mainContainer.Add(colsLayout)
	}

	resultLabel := widget.NewLabelWithStyle("Draw something and click result button.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	resultButton := widget.NewButton("Detect X or O", func() {
		ShowResult(values, weights, bias, theta, resultLabel)
	})
	clearButton := widget.NewButton("Clear", func() {Clear(&values, buttons, resultLabel)})

	mainContainer.Add(resultLabel)
	mainContainer.Add(resultButton)
	mainContainer.Add(clearButton)

	myWindow.ShowAndRun()
}
