package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type configProfileCmd struct {
	cmd *cobra.Command
}

func newConfigProfileCmd() *configProfileCmd {
	cmd := &cobra.Command{
		Use:     "profile",
		Aliases: []string{"p"},
		Short:   "Config profile management",
		Long:    "Allows specific to management of specific profiles in a cfdns config file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	return &configProfileCmd{cmd: cmd}
}

func (cmd *configProfileCmd) loadSubCommands() {
	cmd.cmd.AddCommand(newConfigProfileListCmd())
}
