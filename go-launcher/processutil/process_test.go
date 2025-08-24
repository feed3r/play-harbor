package processutil

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"testing"
)

func TestProcessWrapperWithRealProcess(t *testing.T) {
	pid := int32(os.Getpid())
	proc, err := process.NewProcess(pid)
	if err != nil {
		t.Fatalf("could not get current process: %v", err)
	}
	w := &processWrapper{p: proc}
	if w.Pid() != pid {
		t.Errorf("expected pid %d, got %d", pid, w.Pid())
	}
	name, err := w.Name()
	if err != nil {
		t.Errorf("expected no error from Name, got %v", err)
	}
	if name == "" {
		t.Errorf("expected non-empty name, got empty string")
	}
	// Test WaitForProcessExit with wrapper
	err = WaitForProcessExit(w)
	if err != nil {
		t.Errorf("expected no error from WaitForProcessExit, got %v", err)
	}
}

// Mock for process.Process
type mockProcess struct {
	pid  int32
	name string
}

func (m *mockProcess) Pid() int32            { return m.pid }
func (m *mockProcess) Name() (string, error) { return m.name, nil }

type mockProcessWrapper struct {
	pid  int32
	name string
}

func (pw *mockProcessWrapper) Pid() int32            { return pw.pid }
func (pw *mockProcessWrapper) Name() (string, error) { return pw.name, nil }

func TestProcessWrapperMethods(t *testing.T) {
	w := &mockProcessWrapper{pid: 321, name: "MockName"}
	if w.Pid() != 321 {
		t.Errorf("expected pid 321, got %d", w.Pid())
	}
	name, err := w.Name()
	if err != nil {
		t.Errorf("expected no error from Name, got %v", err)
	}
	if name != "MockName" {
		t.Errorf("expected name 'MockName', got %q", name)
	}
}

func TestFindGameProcess_Found(t *testing.T) {
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	// Simulate process found
	mp := &mockProcess{pid: 123, name: "Game.exe"}
	ProcessesFunc = func() ([]ProcessLike, error) {
		return []ProcessLike{mp}, nil
	}

	proc, err := FindGameProcess("Game.exe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if proc.Pid() != 123 {
		t.Errorf("unexpected pid: %d", proc.Pid())
	}
}

func TestFindGameProcess_ErrorFromProcessesFunc(t *testing.T) {
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	// Simulate error from ProcessesFunc
	ProcessesFunc = func() ([]ProcessLike, error) {
		return nil, fmt.Errorf("mock error")
	}

	_, err := FindGameProcess("Game.exe")
	if err == nil || err.Error() != "Error listing processes: mock error" {
		t.Errorf("expected error from ProcessesFunc, got: %v", err)
	}
}

func TestWaitForProcessExit_ExitsImmediately(t *testing.T) {
	oldPidExists := PidExists
	defer func() { PidExists = oldPidExists }()

	// Simulate process does not exist
	PidExists = func(pid int32) (bool, error) {
		return false, nil
	}

	proc := &process.Process{Pid: 42}
	w := &processWrapper{p: proc}
	err := WaitForProcessExit(w)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
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
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if calls < 3 {
		t.Errorf("expected at least 3 calls, got: %d", calls)
	}
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
	if err != nil {
		t.Errorf("WaitForProcessExit should ignore error, got: %v", err)
	}
}

type errorNameProcess struct {
	pid int32
}

func (e *errorNameProcess) Pid() int32            { return e.pid }
func (e *errorNameProcess) Name() (string, error) { return "", fmt.Errorf("name error") }

func TestFindGameProcess_NameError(t *testing.T) {
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	// Simulate process with Name() error
	ep := &errorNameProcess{pid: 555}
	ProcessesFunc = func() ([]ProcessLike, error) {
		return []ProcessLike{ep}, nil
	}

	_, err := FindGameProcess("Game.exe")
	if err == nil || err.Error() != "Could not find a process with name: Game.exe" {
		t.Errorf("expected not found error, got: %v", err)
	}
}
