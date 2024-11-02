package windows

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func init() {
	fmt.Println("windows init")
}

func AndWindow(parent *fyne.Window) *fyne.Container {
	vbox := container.NewVBox()
	vbox.Add(widget.NewLabel("And gate implementation with Hebb's rule, perceptron network, and adaline."))

	inputsGrid := container.NewGridWithColumns(3)

	firstInput := widget.NewEntry()
	secondInput := widget.NewEntry()
	resultLabel := widget.NewLabel("Type in the values and click the button to see the result.")
	resultButton := widget.NewButtonWithIcon("Result", theme.HelpIcon(), func() {
		ShowResult(firstInput, secondInput, resultLabel, *parent)
	})

	inputsGrid.Add(firstInput)
	inputsGrid.Add(secondInput)
	inputsGrid.Add(resultButton)
	vbox.Add(inputsGrid)
	vbox.Add(resultLabel)

	return vbox
}

func ShowResult(firstEntry *widget.Entry, secondEntry *widget.Entry, resultLabel *widget.Label, parent fyne.Window) {
	firstInputValue, err := strconv.Atoi(firstEntry.Text)
	if err != nil {
		dialog.ShowError(errors.New("first input is not a number"), parent)
	}

	secondInputValue, err := strconv.Atoi(secondEntry.Text)
	if err != nil {
		dialog.ShowError(errors.New("second input is not a number"), parent)
	}

	fmt.Println(firstInputValue)
	fmt.Println(secondInputValue)
}
