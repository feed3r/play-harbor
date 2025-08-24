package launcher

import (
	"os/exec"
	"testing"
)

func TestLaunchGame_CommandMock(t *testing.T) {
	oldExec := ExecCommand
	defer func() { ExecCommand = oldExec }()

	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		if name != "rundll32" {
			t.Errorf("expected rundll32, got %s", name)
		}
		return exec.Command("true") // command that does not fail
	}

	err := LaunchGame("mock-url")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
