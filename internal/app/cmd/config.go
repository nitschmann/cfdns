package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type configCmd struct {
	cmd *cobra.Command
}

func newConfigCmd() *configCmd {
	cmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"c"},
		Short:   "cfdns config file and profile management",
		Long:    "Allows to create and update the config file and its' profiles.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	return &configCmd{cmd: cmd}
}

func (cmd *configCmd) loadSubCommands() {
	cmd.cmd.AddCommand(newConfigCreateCmd())
	cmd.cmd.AddCommand(newConfigFilepathCmd())

	// profile command
	profileCmd := newConfigProfileCmd()
	profileCmd.loadSubCommands()
	cmd.cmd.AddCommand(profileCmd.cmd)
}
