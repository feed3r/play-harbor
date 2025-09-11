package main

import (
	"fmt"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/gamemanager"
	"github.com/feed3r/play-harbor/go-launcher/gui"
)

// main entrypoint
func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("ERROR loading config:", err)
		return
	}

	gm := gamemanager.NewGameManager(cfg)

	gm.FillGameDescriptors()

	// Prepare a list of game names for the GUI
	var gameNames []string
	for _, game := range gm.Games {
		gameNames = append(gameNames, game.DisplayName)
	}

	// Show the GUI window with the games list
	gui.ShowGamesWindow(gameNames)

}
