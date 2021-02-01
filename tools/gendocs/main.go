package main

import (
	"log"

	"github.com/nitschmann/cfdns/internal/app/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cfdnsCmd := cmd.NewRootCmd()
	cfdnsCmd.LoadSubCommands()
	err := doc.GenMarkdownTree(cfdnsCmd.Cmd, "docs/cli")
	if err != nil {
		log.Fatal(err)
	}
}
