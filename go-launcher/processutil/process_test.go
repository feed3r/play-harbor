// Unit tests for processutil package. These tests cover process abstraction, process search, and process exit logic.
package processutil

import (
	"fmt"
	"os"
	"testing"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessWrapperWithRealProcess(t *testing.T) {
	pid := int32(os.Getpid())
	proc, err := process.NewProcess(pid)
	require.NoError(t, err, "could not get current process")
	w := &processWrapper{p: proc}
	assert.Equal(t, pid, w.Pid(), "expected pid uguale")
	name, err := w.Name()
	assert.NoError(t, err, "expected nessun errore da Name")
	assert.NotEmpty(t, name, "expected nome non vuoto")
}

// Mock for process.Process
type mockProcess struct {
	// mockProcess simulates a process for testing purposes
	pid  int32
	name string
}

func (m *mockProcess) Pid() int32            { return m.pid }
func (m *mockProcess) Name() (string, error) { return m.name, nil }

type mockProcessWrapper struct {
	// mockProcessWrapper simulates a processWrapper for testing purposes
	pid  int32
	name string
}

func (pw *mockProcessWrapper) Pid() int32            { return pw.pid }
func (pw *mockProcessWrapper) Name() (string, error) { return pw.name, nil }

func TestProcessWrapperMethods(t *testing.T) {
	w := &mockProcessWrapper{pid: 321, name: "MockName"}
	assert.Equal(t, int32(321), w.Pid(), "expected pid 321")
	name, err := w.Name()
	assert.NoError(t, err, "expected nessun errore da Name")
	assert.Equal(t, "MockName", name, "expected nome 'MockName'")
}

func TestFindGameProcess_Found(t *testing.T) {
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	mp := &mockProcess{pid: 123, name: "Game.exe"}
	ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
		return []ProcessLike{mp}, nil
	}

	proc, err := FindExecutableProcess("Game.exe")
	require.NoError(t, err, "errore inatteso")
	assert.Equal(t, int32(123), proc.Pid(), "pid inatteso")
}

func TestFindGameProcess_ErrorFromProcessesFunc(t *testing.T) {
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
		return nil, fmt.Errorf("mock error")
	}

	_, err := FindExecutableProcess("Game.exe")
	require.Error(t, err)
	assert.Equal(t, "error listing processes: mock error", err.Error(), "errore atteso da ProcessesFunc")
}

func TestWaitForProcessExit_ExitsImmediately(t *testing.T) {
	oldPidExists := PidExists
	defer func() { PidExists = oldPidExists }()

	PidExists = func(pid int32) (bool, error) {
		return false, nil
	}

	proc := &process.Process{Pid: 42}
	w := &processWrapper{p: proc}
	err := WaitForProcessExit(w)
	assert.NoError(t, err, "atteso nessun errore")
}

func TestWaitForProcessExit_WaitsThenExits(t *testing.T) {
	oldPidExists := PidExists
	defer func() { PidExists = oldPidExists }()
	calls := 0
	PidExists = func(pid int32) (bool, error) {
		calls++
		if calls < 3 {
			return true, nil
		}
		return false, nil
	}

	proc := &process.Process{Pid: 99}
	w := &processWrapper{p: proc}
	err := WaitForProcessExit(w)
	assert.NoError(t, err, "atteso nessun errore")
	assert.GreaterOrEqual(t, calls, 3, "attese almeno 3 chiamate")
}

func TestWaitForProcessExit_ErrorFromPidExists(t *testing.T) {
	oldPidExists := PidExists
	defer func() { PidExists = oldPidExists }()
	PidExists = func(pid int32) (bool, error) {
		return false, fmt.Errorf("mock pidexists error")
	}

	proc := &process.Process{Pid: 77}
	w := &processWrapper{p: proc}
	err := WaitForProcessExit(w)
	assert.NoError(t, err, "WaitForProcessExit dovrebbe ignorare l'errore")
}

type errorNameProcess struct {
	// errorNameProcess simulates a process that returns an error on Name()
	pid int32
}

func (e *errorNameProcess) Pid() int32            { return e.pid }
func (e *errorNameProcess) Name() (string, error) { return "", fmt.Errorf("name error") }

func TestFindGameProcess_NameError(t *testing.T) {
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	ep := &errorNameProcess{pid: 555}
	ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
		return []ProcessLike{ep}, nil
	}

	_, err := FindExecutableProcess("Game.exe")
	require.Error(t, err)
	assert.Equal(t, "could not find a process with name: Game.exe", err.Error(), "errore atteso da Name()")
}
