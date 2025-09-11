package main

import (
	"fmt"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/gamemanager"
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

	fmt.Println("Discovered games:")
	for _, game := range gm.Games {
		fmt.Printf("- %s (%s) [%s]\n", game.DisplayName, game.EpicUrl, game.ExeName)
	}

	// r := runlauncher.NewRunLauncher(cfg)
	// err = r.Launch(os.Args[1:])
	// if err != nil {
	// 	fmt.Println("ERROR:", err)
	// 	fmt.Println("Usage: playdock.exe <epicUrl> <exeName>")
	// }

}
