package launcher

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLaunchGame_CommandMock(t *testing.T) {
	oldExec := ExecCommand
	defer func() { ExecCommand = oldExec }()

	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		assert.Equal(t, "rundll32", name, "expected rundll32")
		return exec.Command("true") // command that does not fail
	}

	err := LaunchGame("mock-url")
	assert.NoError(t, err, "expected nessun errore")
}
