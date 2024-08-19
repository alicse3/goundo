package main

import (
	"github.com/alicse3/goundo/cmd"
	"github.com/alicse3/goundo/internal/config"
)

func init() {
	config.InitSetup()
}

func main() {
	cmd.Execute()
}
