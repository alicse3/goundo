package cmd

import (
	"fmt"
)

const help = `usage: goundo <command> [arguments]

	The most commonly used goundo commands are:
		version     Show the app's version number
		help        Show this help message
		configure   Configure the app settings`

// helpHandler handles the help command
func helpHandler() {
	fmt.Println(help)
}
