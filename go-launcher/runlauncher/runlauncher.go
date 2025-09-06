package runlauncher

import (
	"fmt"
	"strings"
	"time"

	"github.com/feed3r/play-harbor/go-launcher/args"
	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/launcher"
	"github.com/feed3r/play-harbor/go-launcher/processutil"
)

type RunLauncher struct {
	Config              *config.Config
	SleepFunc           func()
	PollGameProcessFunc func(string) (processutil.ProcessLike, error)
}

func NewRunLauncher(cfg *config.Config) *RunLauncher {
	return &RunLauncher{
		Config: cfg,
		SleepFunc: func() {
			time.Sleep(cfg.Global.SleepWithoutManager)
		},
	}
}

// Global function variables to allow mocking in tests
var LaunchGameFunc = launcher.LaunchGame
var WaitForProcessExitFunc = processutil.WaitForProcessExit

// Check if the game manager platform client (IE: the Epic Games Launcher)
// is active
var IsManagerRunning = func(executableName string) (bool, error) {
	procs, err := processutil.FindExecutableProcess(executableName)
	if err != nil {
		if strings.Contains(err.Error(), "could not find a process with name") {
			return false, nil // Not running, but not a real error
		}
		return false, err
	}
	return procs != nil, nil
}

// Orchestrator for launching the game
func (r *RunLauncher) Launch(argsList []string) error {
	epicUrl, exeName, err := args.ParseArgs(argsList)
	if err != nil {
		return err
	}

	// Check if the game manager is running
	managerRunning, err := IsManagerRunning(r.Config.EpicGamesStore.Executable)
	if err != nil {
		return fmt.Errorf("error checking for game manager process: %v", err)
	}

	if managerRunning {
		fmt.Println("Game manager detected, using shorter wait time")
		r.SleepFunc = func() {
			time.Sleep(r.Config.Global.SleepWithManager)
		}
	} else {
		fmt.Println("No game manager detected, using longer wait time")
		// SleepFunc remains as default
	}

	if err := LaunchGameFunc(epicUrl); err != nil {
		return fmt.Errorf("failed to start URL: %v", err)
	}

	// Wait a few seconds to give the game time to start
	// You can modularize this in the future
	r.SleepFunc()

	// Find the game process with a polling
	proc, err := r.PollGameProcessFunc(exeName)
	if err != nil {
		return err
	}

	return WaitForProcessExitFunc(proc)
}

// Find the game process executing. Poll actively for the process
// for a maximum amount of time
func (r *RunLauncher) PollGameProcess(exeName string) (processutil.ProcessLike, error) {
	pollingTrialCounter := r.Config.Global.MaxPollingAttempts

	for pollingTrialCounter > 0 {
		pollingTrialCounter--
		proc, err := processutil.FindExecutableProcess(exeName)
		if err != nil {
			return nil, err
		} else if proc != nil {
			fmt.Printf("Found game process with PID %d\n", proc.Pid())
			return proc, nil
		}
		time.Sleep(r.Config.Global.PollingInterval)
	}
	return nil, fmt.Errorf("game process not found after %d attempts", r.Config.Global.MaxPollingAttempts)
}
