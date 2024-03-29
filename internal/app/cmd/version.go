package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Version of cfdns",
		Long:    "The version of the cfdns tool itself",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(AppVersion)
		},
	}

	return cmd
}
