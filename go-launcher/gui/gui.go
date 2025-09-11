package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowGamesWindow displays a window with the list of games using Fyne.
func ShowGamesWindow(games []string) {
	a := app.New()
	w := a.NewWindow("Games List")
	w.Resize(fyne.NewSize(400, 600))

	label := widget.NewLabel("Installed Games:")
	list := widget.NewList(
		func() int { return len(games) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(games[i])
		},
	)

	content := container.NewVBox(label, list)
	w.SetContent(content)
	w.ShowAndRun()
}
