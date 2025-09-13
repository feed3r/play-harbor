package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/feed3r/play-harbor/go-launcher/gamemanager"
	"github.com/feed3r/play-harbor/go-launcher/runlauncher"
)

// ShowGamesWindow displays a window with the list of games using Fyne.
func ShowGamesWindow(runlauncher *runlauncher.RunLauncher, games []*gamemanager.GameDescriptor) {
	a := app.New()
	w := a.NewWindow("Games List")
	w.Resize(fyne.NewSize(400, 600))

	label := widget.NewLabel("Installed Games:")

	list := widget.NewList(
		func() int { return len(games) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(games[i].DisplayName)
		},
	)

	// Add a Close button below the games list
	closeBtn := widget.NewButton("Close", func() {
		w.Close()
	})

	// Helper function to launch a game (replace with your logic)
	launchGame := func(game *gamemanager.GameDescriptor) {
		argsList := []string{game.EpicUrl, game.ExeName} // or whatever args you need
		err := runlauncher.Launch(argsList)
		if err != nil {
			// Handle error (e.g., show a popup)
			errorPopup := fyne.CurrentApp().NewWindow("Error")
			errorMsg := widget.NewLabel("Failed to launch game: " + err.Error())
			okBtn := widget.NewButton("Ok", func() {
				errorPopup.Close()
			})
			content := container.NewVBox(errorMsg, okBtn)
			errorPopup.SetContent(content)
			errorPopup.Resize(fyne.NewSize(300, 100))
			errorPopup.Show()
			errorPopup.Canvas().Focus(okBtn)
			return
		}

		// Show a popup window with the launch message and an Ok button
		popup := fyne.CurrentApp().NewWindow("Launching Game")
		msg := widget.NewLabel("You are launching the game: " + game.DisplayName + "\nExecutable: " + game.ExeName + "\nURL: " + game.EpicUrl)
		okBtn := widget.NewButton("Ok", func() {
			popup.Close()
			w.Canvas().Focus(list)
		})
		content := container.NewVBox(msg, okBtn)
		popup.SetContent(content)
		popup.Resize(fyne.NewSize(300, 100))
		popup.Show()
		popup.Canvas().Focus(okBtn)
	}

	list.OnSelected = func(id int) {
		launchGame(games[id])
		list.Unselect(id)
	}

	// Use container.NewBorder to put the Close button at the bottom and let the list expand
	content := container.NewBorder(label, closeBtn, nil, nil, list)
	w.SetContent(content)
	w.Canvas().Focus(list)
	w.ShowAndRun()
}
