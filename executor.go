package resticgo

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
)

func execute(command []string) (string, error) {
	slog.Debug(fmt.Sprintf("Executing restic command: %s", command))
	if len(command) < 2 {
		return "", errors.New("invalid restic command")
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		var execError *exec.Error
		if errors.As(err, &execError) {
			return "", &NoResticBinaryError{"Cannot find restic installed"}
		}
		slog.Debug(fmt.Sprintf("Restic command completed with return code %d", cmd.ProcessState.ExitCode()))
		return "", &ResticFailedError{err: fmt.Sprintf("Restic failed with exit code %d: %s", cmd.ProcessState.ExitCode(), stderr.String()), ExitCode: cmd.ProcessState.ExitCode()}
	}

	slog.Debug(fmt.Sprintf("Restic command completed with return code %d", cmd.ProcessState.ExitCode()))
	return stdout.String(), nil
}
