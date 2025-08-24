package launcher

import "os/exec"

var ExecCommand = exec.Command

func LaunchGame(epicUrl string) error {
	cmd := ExecCommand("rundll32", "url.dll,FileProtocolHandler", epicUrl)
	return cmd.Start()
}
