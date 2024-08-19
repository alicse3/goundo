package cmd

import (
	"flag"
	"os"
)

func versionHandler() {
	// Create a new flag set
	flagSet := flag.NewFlagSet("version", flag.ExitOnError)

	// Define custom usage function
	flagSet.Usage = func() {
		println("Usage: goundo version [flags]")
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
		println("1.0.0")
	}
}
