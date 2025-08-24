package main

import (
	"fmt"
	"github.com/feed3r/play-harbor/go-launcher/runlauncher"
	"os"
)

// main entrypoint
func main() {
	err := runlauncher.RunLauncher(os.Args[1:])
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("Usage: playdock.exe <epicUrl> <exeName>")
	}
}
