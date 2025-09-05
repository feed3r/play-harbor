package processutil

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
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

var ProcessesFunc = func(searchName ...string) ([]ProcessLike, error) {
	procs, err := Processes()
	if err != nil {
		return nil, err
	}

	// If no search criteria, return all processes
	if len(searchName) == 0 {
		out := make([]ProcessLike, len(procs))
		for i, p := range procs {
			out[i] = &processWrapper{p: p}
		}
		return out, nil
	}

	// Search for specific process(es)
	search := strings.ToLower(searchName[0])
	searchWithExt := search
	if !strings.HasSuffix(search, ".exe") {
		searchWithExt = search + ".exe"
	}

	var out []ProcessLike
	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue // Skip processes we can't get names for
		}

		lowerName := strings.ToLower(name)
		// Check both with and without .exe extension
		if lowerName == search || lowerName == searchWithExt {
			out = append(out, &processWrapper{p: p})
		} else if strings.HasSuffix(lowerName, ".exe") {
			nameWithoutExt := strings.TrimSuffix(lowerName, ".exe")
			if nameWithoutExt == search {
				out = append(out, &processWrapper{p: p})
			}
		}
	}

	return out, nil
}

var Processes = process.Processes
var PidExists = process.PidExists

func FindGameProcess(exeName string) (ProcessLike, error) {
	procs, err := ProcessesFunc(exeName)
	if err != nil {
		return nil, fmt.Errorf("error listing processes: %v", err)
	}

	// Return the first process with a valid name
	for _, proc := range procs {
		if _, err := proc.Name(); err == nil {
			return proc, nil
		}
	}

	return nil, fmt.Errorf("could not find a process with name: %s", exeName)
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
