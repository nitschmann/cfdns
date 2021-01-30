package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nitschmann/cfdns/internal/pkg/config"
)

func newConfigFilepathCmd() *cobra.Command {
	cmd := &cobra.Command{
		Aliases: []string{"file"},
		Use:     "filepath",
		Short:   "Filepath of the auto-loaded config file",
		Long:    "Filepath of the auto-loaded config file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config.AutoFilepath())
		},
	}

	return cmd
}
