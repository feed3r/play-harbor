package runlauncher

import (
	"errors"
	"testing"
	"time"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/feed3r/play-harbor/go-launcher/processutil"
	"github.com/stretchr/testify/assert"
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
	WaitForProcessExitFunc = func(proc processutil.ProcessLike) error {
		return mockWaitForProcessExit(proc)
	}
}

// Helper per creare un RunLauncher di default per i test
func newTestRunLauncher() *RunLauncher {
	rl := &RunLauncher{
		Config: &config.Config{
			Global: config.GlobalConfig{
				SleepWithManager:    1 * time.Millisecond,
				SleepWithoutManager: 1 * time.Millisecond,
				MaxPollingAttempts:  2,
				PollingInterval:     1 * time.Millisecond,
			},
		},
	}
	rl.SleepFunc = func() {}
	rl.PollGameProcessFunc = func(name string) (processutil.ProcessLike, error) {
		return &mockProcess{}, nil
	}
	return rl
}

func TestRunLauncher_PollGameProcess_RetriesOnNotFound(t *testing.T) {
	oldProcessesFunc := processutil.ProcessesFunc
	defer func() { processutil.ProcessesFunc = oldProcessesFunc }()

	calls := 0
	processutil.ProcessesFunc = func(searchName ...string) ([]processutil.ProcessLike, error) {
		calls++
		if calls < 3 {
			return nil, nil // not found yet
		}
		return []processutil.ProcessLike{&mockProcess{}}, nil
	}

	r := newTestRunLauncher()
	r.Config.Global.MaxPollingAttempts = 5
	r.Config.Global.PollingInterval = 1 * time.Millisecond

	proc, err := r.PollGameProcess("game.exe")
	assert.NoError(t, err, "a process not found yet should not abort polling")
	assert.NotNil(t, proc)
	assert.GreaterOrEqual(t, calls, 3, "should retry until the process appears")
}

func TestRunLauncher_PollGameProcess_PropagatesRealError(t *testing.T) {
	oldProcessesFunc := processutil.ProcessesFunc
	defer func() { processutil.ProcessesFunc = oldProcessesFunc }()

	calls := 0
	processutil.ProcessesFunc = func(searchName ...string) ([]processutil.ProcessLike, error) {
		calls++
		return nil, errors.New("boom")
	}

	r := newTestRunLauncher()
	r.Config.Global.MaxPollingAttempts = 5
	r.Config.Global.PollingInterval = 1 * time.Millisecond

	_, err := r.PollGameProcess("game.exe")
	assert.Error(t, err, "a real enumeration error must not be confused with 'process not found'")
	assert.Equal(t, 1, calls, "a real error must abort polling immediately")
}

func TestRunLauncher_ManagerRunning(t *testing.T) {
	r := newTestRunLauncher()
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

	err := r.Launch([]string{"epic://game", "game.exe"})
	assert.NoError(t, err, "RunLauncher con manager attivo dovrebbe restituire nil")
}

func TestRunLauncher_ManagerNotRunning(t *testing.T) {
	r := newTestRunLauncher()
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

	err := r.Launch([]string{"epic://game", "game.exe"})
	assert.NoError(t, err, "RunLauncher senza manager dovrebbe restituire nil")
}

func TestRunLauncher_LaunchGameError(t *testing.T) {
	r := newTestRunLauncher()
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

	err := r.Launch([]string{"epic://game", "game.exe"})
	assert.Error(t, err, "RunLauncher dovrebbe restituire errore se LaunchGame fallisce")
}
