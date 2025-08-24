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
	for _, p := range procs {
		name, _ := p.Name()
		if strings.EqualFold(name, exeName+".exe") || strings.EqualFold(name, exeName) {
			return p, nil
		}
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
