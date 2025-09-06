// Unit tests for processutil package. These tests cover process abstraction, process search, and process exit logic.
package processutil

import (
	"fmt"
	"os"
	"testing"

	"github.com/shirou/gopsutil/v3/process"
)

func TestProcessWrapperWithRealProcess(t *testing.T) {
	// Test processWrapper with the real current process
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
	// Test the mock process wrapper methods for correct values
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
	// Test FindGameProcess when the process is found
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	// Simulate process found
	mp := &mockProcess{pid: 123, name: "Game.exe"}
	ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
		return []ProcessLike{mp}, nil
	}

	proc, err := FindExecutableProcess("Game.exe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if proc.Pid() != 123 {
		t.Errorf("unexpected pid: %d", proc.Pid())
	}
}

func TestFindGameProcess_ErrorFromProcessesFunc(t *testing.T) {
	// Test FindGameProcess when ProcessesFunc returns an error
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	// Simulate error from ProcessesFunc
	ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
		return nil, fmt.Errorf("mock error")
	}

	_, err := FindExecutableProcess("Game.exe")
	if err == nil || err.Error() != "error listing processes: mock error" {
		t.Errorf("expected error from ProcessesFunc, got: %v", err)
	}
}

func TestWaitForProcessExit_ExitsImmediately(t *testing.T) {
	// Test WaitForProcessExit when the process does not exist (should exit immediately)
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
	// Test WaitForProcessExit when the process exists for a few cycles before exiting
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
	// Test WaitForProcessExit when PidExists returns an error (should ignore the error)
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
	// errorNameProcess simulates a process that returns an error on Name()
	pid int32
}

func (e *errorNameProcess) Pid() int32            { return e.pid }
func (e *errorNameProcess) Name() (string, error) { return "", fmt.Errorf("name error") }

func TestFindGameProcess_NameError(t *testing.T) {
	// Test FindGameProcess when the process returns an error on Name()
	oldFunc := ProcessesFunc
	defer func() { ProcessesFunc = oldFunc }()

	// Simulate process with Name() error
	ep := &errorNameProcess{pid: 555}
	ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
		return []ProcessLike{ep}, nil
	}

	_, err := FindExecutableProcess("Game.exe")
	if err == nil || err.Error() != "could not find a process with name: Game.exe" {
		t.Errorf("expected not found error, got: %v", err)
	}
}
