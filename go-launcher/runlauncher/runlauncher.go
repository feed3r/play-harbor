package runlauncher

import (
	"fmt"
	"github.com/feed3r/play-harbor/go-launcher/args"
	"github.com/feed3r/play-harbor/go-launcher/launcher"
	"github.com/feed3r/play-harbor/go-launcher/processutil"
)

// Orchestrator for launching the game
func RunLauncher(argsList []string) error {
	epicUrl, exeName, err := args.ParseArgs(argsList)
	if err != nil {
		return err
	}
	if err := launcher.LaunchGame(epicUrl); err != nil {
		return fmt.Errorf("Failed to start URL: %v", err)
	}
	// Wait a few seconds to give the game time to start
	// You can modularize this in the future
	// time.Sleep(5 * time.Second)
	proc, err := processutil.FindGameProcess(exeName)
	if err != nil {
		return err
	}
	return processutil.WaitForProcessExit(proc)
}
