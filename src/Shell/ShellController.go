package Shell

import (
	"io"
	"os/exec"
)

func Execute(cmd string, arguments []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	command := exec.Cmd{
		Path:   cmd,
		Args:   arguments,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	return command.Run()
}
