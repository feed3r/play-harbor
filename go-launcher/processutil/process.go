package processutil

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"strings"
	"time"
)

// ProcessLike abstracts process.Process for testability
type ProcessLike interface {
	Pid() int32
	Name() (string, error)
}

// processWrapper wraps *process.Process to implement ProcessLike
type processWrapper struct {
	p *process.Process
}

func (pw *processWrapper) Pid() int32 {
	return pw.p.Pid
}
func (pw *processWrapper) Name() (string, error) {
	return pw.p.Name()
}

var ProcessesFunc = func() ([]ProcessLike, error) {
	procs, err := Processes()
	if err != nil {
		return nil, err
	}
	out := make([]ProcessLike, len(procs))
	for i, p := range procs {
		out[i] = &processWrapper{p: p}
	}
	return out, nil
}

var Processes = process.Processes
var PidExists = process.PidExists

func FindGameProcess(exeName string) (ProcessLike, error) {
	procs, err := ProcessesFunc()
	if err != nil {
		return nil, fmt.Errorf("Error listing processes: %v", err)
	}
	
	// Build a hashtable for O(1) lookup
	processMap := make(map[string]ProcessLike)
	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue // Skip processes we can't get names for
		}
		// Store both with and without .exe extension for flexibility
		lowerName := strings.ToLower(name)
		processMap[lowerName] = p
		if strings.HasSuffix(lowerName, ".exe") {
			nameWithoutExt := strings.TrimSuffix(lowerName, ".exe")
			processMap[nameWithoutExt] = p
		}
	}
	
	// Try to find the process with O(1) lookup
	searchName := strings.ToLower(exeName)
	if proc, found := processMap[searchName]; found {
		return proc, nil
	}
	if proc, found := processMap[searchName+".exe"]; found {
		return proc, nil
	}
	
	return nil, fmt.Errorf("Could not find a process with name: %s", exeName)
}

func WaitForProcessExit(proc ProcessLike) error {
	for {
		exists, _ := PidExists(proc.Pid())
		if !exists {
			break
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
