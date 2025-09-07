package main

import (
	"fmt"
	"os"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/runlauncher"
)

// main entrypoint
func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("ERROR loading config:", err)
		return
	}

	r := runlauncher.NewRunLauncher(cfg)
	err = r.Launch(os.Args[1:])
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("Usage: playdock.exe <epicUrl> <exeName>")
	}
}
