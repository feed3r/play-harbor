package main

import (
	"fmt"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/gamemanager"
	"github.com/feed3r/play-harbor/go-launcher/gui"
	"github.com/feed3r/play-harbor/go-launcher/runlauncher"
)

// main entrypoint
func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("ERROR loading config:", err)
		return
	}

	gm := gamemanager.NewGameManager(cfg)
	rl := runlauncher.NewRunLauncher(cfg)

	gm.FillGameDescriptors()

	// Show the GUI window with the games list
	gui.ShowGamesWindow(rl, gm.Games)

}
