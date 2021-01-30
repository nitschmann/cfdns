package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	configServ "github.com/nitschmann/cfdns/internal/app/service/config"
	"github.com/nitschmann/cfdns/internal/pkg/config"
)

func newConfigProfileListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "names"},
		Short:   "List the currently defined profile names",
		Long:    "List the currently defined profile names in the given config file.",
		Run: func(cmd *cobra.Command, args []string) {
			var configFilepath string

			configFilepath, err := cmd.Flags().GetString("file")
			if err != nil {
				printCliErrorAndExit(err)
			}

			if configFilepath == "" {
				configFilepath = config.AutoFilepath()
			}

			profileListService := configServ.NewProfileListService(configFilepath)
			list, err := profileListService.Get()
			if err != nil {
				printCliErrorAndExit(err)
			}

			for name, _ := range list {
				fmt.Println(name)
			}
		},
	}

	cmd.Flags().StringP("file", "f", "", "The config file to use (If not given the autoloaded file will be used)")

	return cmd
}
