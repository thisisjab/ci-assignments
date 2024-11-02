package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"xo-detection/ui/windows"
)

func StartUI() {
	mainApp := app.New()
	mainWindow := mainApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), windows.HomeWindow()),
		container.NewTabItemWithIcon("And Gate", theme.HelpIcon(), windows.AndWindow(&mainWindow)),
		container.NewTabItemWithIcon("XO Detection", theme.GridIcon(), windows.XOWindow(&mainWindow)),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	mainWindow.SetContent(tabs)
	mainWindow.Resize(fyne.NewSize(700, 450))
	mainWindow.ShowAndRun()
}
