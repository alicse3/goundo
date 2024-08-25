package cmd

import (
	"fmt"
	"os"
	"strings"
)

// HandleCommands handles the command line arguments and calls the appropriate handler function
// based on the provided command.
func HandleCommands() {
	// If there are no arguments, show the help message and exit
	if len(os.Args) < 2 {
		helpHandler()
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
		backupHandler()
	case "list":
		listHandler()
	default:
		fmt.Println("Unknown command:", command)
		helpHandler()
	}
}
