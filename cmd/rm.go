package cmd

import (
	"os"
	"os/exec"
)

func rmHandler() {
	// Get rm args
	args := os.Args[2:]

	// TODO: Parse rm args
	// TODO: Backup data

	// Structure command
	cmd := exec.Command("rm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run command
	if err := cmd.Run(); err != nil {
		println("Error is:", err.Error())
	} else {
		println("Success")
	}

	// TODO: Undo backup in case of failures
	// TODO: Track info in DB in case of success
}
