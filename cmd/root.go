package cmd

import (
	"os"
	"strings"
)

// HandleCommands handles the command line arguments and calls the appropriate handler function
// based on the provided command.
func HandleCommands() {
	// If there are no arguments, show the help message and exit
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	// Get the command and call command handler
	command := strings.TrimSpace(os.Args[1])

	// Handle command
	switch command {
	case "configure":
		configHandler()
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
