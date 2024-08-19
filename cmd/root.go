package cmd

import (
	"os"
)

func Execute() {
	// If there are no arguments, show the help message and exit
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	// Get the command and call command handler
	command := os.Args[1]
	switch command {
	case "version":
		versionHandler()
	case "help":
		helpHandler()
	case "rm":
		rmHandler()
	default:
		println("Unknown command")
		showHelp()
	}
}
