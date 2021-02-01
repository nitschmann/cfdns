package cmd

import (
	"github.com/spf13/cobra"

	configServ "github.com/nitschmann/cfdns/internal/app/service/config"
	"github.com/nitschmann/cfdns/internal/pkg/config"
)

func newConfigProfileDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [NAME]",
		Aliases: []string{"d"},
		Short:   "Delete a config profile",
		Long:    "Delete a config profile with the given name",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var configFilepath string

			configFilepath, err := cmd.Flags().GetString("file")
			if err != nil {
				printCliErrorAndExit(err)
			}

			if configFilepath == "" {
				configFilepath = config.AutoFilepath()
			}

			profileDeleteService := configServ.NewProfileDeleteService(configFilepath)
			err = profileDeleteService.DeleteProfile(args[0])
			if err != nil {
				printCliErrorAndExit(err)
			}
		},
	}

	cmd.Flags().StringP("file", "f", "", "The config file to use (If not given the autoloaded config file will be used)")

	return cmd
}
