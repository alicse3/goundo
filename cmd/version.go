package cmd

import "fmt"

const latestVersion = "v1.0.0"

// versionHandler prints the app's latest version
func versionHandler() {
	fmt.Println(latestVersion)
}
