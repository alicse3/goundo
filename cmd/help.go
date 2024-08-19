package cmd

import (
	"flag"
	"os"
)

func showHelp() {
	help := `usage: goundo <command> [arguments]

	The most commonly used goundo commands are:
		version  Print the version number of goundo
		help     Print this help message

	Use "goundo <command> -h" for more information about a command.`

	println(help)
}

func helpHandler() {
	// Create a new flag set
	flagSet := flag.NewFlagSet("help", flag.ExitOnError)

	// Define custom usage function
	flagSet.Usage = func() {
		println("Usage: goundo help [flags]")
		println("Flags:")
		flagSet.PrintDefaults()
	}

	// Define the help flag
	helpFlg := flagSet.Bool("h", false, "Display help information")

	// Parse the command-line arguments
	if err := flagSet.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	// Check if the help flag is set
	if *helpFlg {
		flagSet.Usage()
	} else {
		showHelp()
	}
}
