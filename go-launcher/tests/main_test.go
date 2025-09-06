package main

import (
	"errors"
	"testing"
	"time"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/processutil"
	"github.com/feed3r/play-harbor/go-launcher/runlauncher"
)

// Mock process
type mockProcess struct{}

func (m *mockProcess) Pid() int32            { return 42 }
func (m *mockProcess) Name() (string, error) { return "Game.exe", nil }

func newTestRunLauncher() *runlauncher.RunLauncher {
	cfg := &config.Config{
		Global: config.GlobalConfig{
			SleepWithManager:    1 * time.Millisecond,
			SleepWithoutManager: 1 * time.Millisecond,
			MaxPollingAttempts:  2,
			PollingInterval:     1 * time.Millisecond,
		},
		EpicGamesStore: config.EpicGamesStoreConfig{
			Executable: "EpicGamesLauncher.exe",
		},
	}
	r := runlauncher.NewRunLauncher(cfg)
	r.PollGameProcessFunc = func(name string) (processutil.ProcessLike, error) {
		return &mockProcess{}, nil
	}
	runlauncher.IsManagerRunning = func(executableName string) (bool, error) {
		return false, nil
	}
	runlauncher.LaunchGameFunc = func(url string) error {
		return nil
	}
	runlauncher.WaitForProcessExitFunc = func(proc processutil.ProcessLike) error {
		return nil
	}
	return r
}

func TestRunLauncher_MissingArgs(t *testing.T) {
	r := newTestRunLauncher()
	err := r.Launch([]string{})
	if err == nil {
		t.Error("Expected error for missing args, got nil")
	}
}

func TestRunLauncher_CommandError(t *testing.T) {
	r := newTestRunLauncher()
	runlauncher.LaunchGameFunc = func(url string) error {
		return errors.New("mock error")
	}
	err := r.Launch([]string{"mock-url", "Game.exe"})
	if err == nil {
		t.Error("Expected error for command error, got nil")
	}
}

func TestRunLauncher_ProcessNotFound(t *testing.T) {
	r := newTestRunLauncher()
	r.PollGameProcessFunc = func(name string) (processutil.ProcessLike, error) {
		return nil, errors.New("process not found")
	}
	err := r.Launch([]string{"mock-url", "Game.exe"})
	if err == nil {
		t.Error("Expected error for process not found, got nil")
	}
}

func TestRunLauncher_ProcessFound(t *testing.T) {
	r := newTestRunLauncher()
	err := r.Launch([]string{"mock-url", "Game.exe"})
	if err != nil {
		t.Errorf("Expected nil error for process found, got %v", err)
	}
}
