package main

import (
	"commands"
	"os"
)

func main() {
	// Process the input file and give and interactive shell
	if len(os.Args) > 1 && "" != os.Args[1] {
		commands.NewFileCmdProcessor(os.Args[1]).Process()
	}
	// Interactive shell will provide to user to operate
	// with file data and console process
	commands.NewShell().Process()
}
