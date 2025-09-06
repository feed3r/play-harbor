package runlauncher

import (
	"errors"
	"testing"

	"github.com/feed3r/play-harbor/go-launcher/processutil"
)

// ProcessLike mock
type mockProcess struct{}

func (m *mockProcess) Pid() int32            { return 1234 }
func (m *mockProcess) Name() (string, error) { return "game.exe", nil }

// Mock delle dipendenze come variabili globali
var (
	mockIsManagerRunning   func(string) (bool, error)
	mockLaunchGame         func(string) error
	mockPollGameProcess    func(string) (processutil.ProcessLike, error)
	mockWaitForProcessExit func(processutil.ProcessLike) error
)

// Override le funzioni globali nel package runlauncher
func init() {
	IsManagerRunning = func(executableName string) (bool, error) {
		return mockIsManagerRunning(executableName)
	}
	LaunchGameFunc = func(url string) error {
		return mockLaunchGame(url)
	}
	PollGameProcessFunc = func(name string) (processutil.ProcessLike, error) {
		return mockPollGameProcess(name)
	}
	WaitForProcessExitFunc = func(proc processutil.ProcessLike) error {
		return mockWaitForProcessExit(proc)
	}
}

func TestRunLauncher_ManagerRunning(t *testing.T) {
	SleepFunc = func() {}
	mockIsManagerRunning = func(executableName string) (bool, error) {
		return true, nil
	}
	mockLaunchGame = func(url string) error {
		return nil
	}
	mockPollGameProcess = func(name string) (processutil.ProcessLike, error) {
		return &mockProcess{}, nil
	}
	mockWaitForProcessExit = func(proc processutil.ProcessLike) error {
		return nil
	}

	err := RunLauncher([]string{"epic://game", "game.exe"})
	if err != nil {
		t.Errorf("RunLauncher con manager attivo dovrebbe restituire nil, ottenuto: %v", err)
	}
}

func TestRunLauncher_ManagerNotRunning(t *testing.T) {
	SleepFunc = func() {}
	mockIsManagerRunning = func(executableName string) (bool, error) {
		return false, nil
	}
	mockLaunchGame = func(url string) error {
		return nil
	}
	mockPollGameProcess = func(name string) (processutil.ProcessLike, error) {
		return &mockProcess{}, nil
	}
	mockWaitForProcessExit = func(proc processutil.ProcessLike) error {
		return nil
	}

	err := RunLauncher([]string{"epic://game", "game.exe"})
	if err != nil {
		t.Errorf("RunLauncher senza manager dovrebbe restituire nil, ottenuto: %v", err)
	}
}

func TestRunLauncher_LaunchGameError(t *testing.T) {
	SleepFunc = func() {}
	mockIsManagerRunning = func(executableName string) (bool, error) {
		return true, nil
	}
	mockLaunchGame = func(url string) error {
		return errors.New("errore lancio")
	}
	mockPollGameProcess = func(name string) (processutil.ProcessLike, error) {
		return &mockProcess{}, nil
	}
	mockWaitForProcessExit = func(proc processutil.ProcessLike) error {
		return nil
	}

	err := RunLauncher([]string{"epic://game", "game.exe"})
	if err == nil {
		t.Error("RunLauncher dovrebbe restituire errore se LaunchGame fallisce")
	}
}
