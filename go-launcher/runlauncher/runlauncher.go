package runlauncher

import (
	"fmt"
	"time"

	"github.com/feed3r/play-harbor/go-launcher/args"
	"github.com/feed3r/play-harbor/go-launcher/launcher"
	"github.com/feed3r/play-harbor/go-launcher/processutil"
)

const SLEEP_WITHOUT_MANAGER = 30 * time.Second
const SLEEP_WITH_MANAGER = 10 * time.Second

const MANAGER_EXECUTABLE_NAME = "EpicGamesLauncher.exe"
const MAX_POLLING_ATTEMPTS = 20
const POLLING_INTERVAL = 1 * time.Second

// Global function variables to allow mocking in tests
var LaunchGameFunc = launcher.LaunchGame
var PollGameProcessFunc = PollGameProcess
var WaitForProcessExitFunc = processutil.WaitForProcessExit

// Check if the game manager platform client (IE: the Epic Games Launcher)
// is active
var IsManagerRunning = func(executableName string) (bool, error) {
	procs, err := processutil.FindExecutableProcess(executableName)
	if err != nil {
		return false, err
	}
	return procs != nil, nil
}

var SleepFunc = func() {
	time.Sleep(SLEEP_WITHOUT_MANAGER)
}

// Orchestrator for launching the game
func RunLauncher(argsList []string) error {
	epicUrl, exeName, err := args.ParseArgs(argsList)
	if err != nil {
		return err
	}

	// Check if the game manager is running
	managerRunning, err := IsManagerRunning(MANAGER_EXECUTABLE_NAME)
	if err != nil {
		return fmt.Errorf("error checking for game manager process: %v", err)
	}

	if managerRunning {
		fmt.Println("Game manager detected, using shorter wait time")
		SleepFunc = func() {
			time.Sleep(SLEEP_WITH_MANAGER)
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
	SleepFunc()

	// Find the game process with a polling
	proc, err := PollGameProcessFunc(exeName)
	if err != nil {
		return err
	}

	return WaitForProcessExitFunc(proc)
}

// Find the game process executing. Poll actively for the process
// for a maximum amount of time
func PollGameProcess(exeName string) (processutil.ProcessLike, error) {
	pollingTrialCounter := MAX_POLLING_ATTEMPTS

	for pollingTrialCounter > 0 {
		pollingTrialCounter--
		proc, err := processutil.FindExecutableProcess(exeName)
		if err != nil {
			return nil, err
		} else if err == nil && proc != nil {
			fmt.Printf("Found game process with PID %d\n", proc.Pid())
			return proc, nil
		}
		time.Sleep(POLLING_INTERVAL)
	}
	return nil, fmt.Errorf("game process not found after %d attempts", MAX_POLLING_ATTEMPTS)
}
