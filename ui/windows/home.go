package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeWindow() *fyne.Container {
	vbox := container.NewVBox()
	vbox.Add(widget.NewLabel("Hello. This software contains implementations for various neural network algorithms."))
	vbox.Add(widget.NewLabel("This whole project includes all assignments for computational intelligence course 2023 by Dr. Koohestani."))
	return vbox
}
