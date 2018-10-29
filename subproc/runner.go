package subproc

import "os/exec"

func RunSubprocess(command string, args []string) ([]byte, error) {
	// TODO: Run with timeout using context
	cmd := exec.Command(command, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	return stdoutStderr, err
}
