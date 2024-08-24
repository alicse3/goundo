package util

import (
	"fmt"
	"os"
)

// ExitOnError checks if an error occurred and exits the program with a non-zero status code.
func ExitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s, err: %v\n", msg, err)
	} else {
		fmt.Println(msg)
	}

	os.Exit(1)
}
