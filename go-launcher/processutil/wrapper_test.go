//go:build test
// +build test

package processutil

// processWrapper mock per test, non accede al filesystem

type processWrapper struct {
	pid  int32
	name string
}

func (pw *processWrapper) Pid() int32 {
	return pw.pid
}
func (pw *processWrapper) Name() (string, error) {
	return pw.name, nil
}
