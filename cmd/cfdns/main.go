package main

import "github.com/nitschmann/cfdns/internal/app/cmd"

// Version is the global version the cfdns CLI tool
var Version string

func main() {
	cmd.AppVersion = Version
	cmd.Execute()
}
