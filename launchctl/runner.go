package launchctl

import (
	"os/exec"
)

const launchtlBinary = "/bin/launchctl"

func Load(taskName string) (string, error) {
	output, error := RunSubprocess(launchtlBinary, []string{"load", taskName})
	return string(output), error
}

func RunSubprocess(command string, args []string) ([]byte, error) {
	// TODO: Run with timeout using context
	cmd := exec.Command(command, args...)
	return cmd.CombinedOutput()
}
