package main

import (
	"github.com/alicse3/goundo/cmd"
)

func init() {
	cmd.InitSetup()
}

func main() {
	cmd.HandleCommands()
}
