package main

import "github.com/nitschmann/cfd/internal/app/cmd"

// Version is the global version the cfd CLI tool
var Version string

func main() {
	cmd.AppVersion = Version
	cmd.Execute()
}
