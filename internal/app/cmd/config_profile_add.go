package cmd

import (
	"github.com/spf13/cobra"

	configServ "github.com/nitschmann/cfdns/internal/app/service/config"
	"github.com/nitschmann/cfdns/internal/pkg/config"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

func newConfigProfileAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [NAME]",
		Aliases: []string{"a"},
		Short:   "Add a new config profile",
		Long:    "Add a new config profile with the given name and specified parameter flags.",
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

			apiKey, err := cmd.Flags().GetString("api-key")
			if err != nil {
				printCliErrorAndExit(err)
			}
			email, err := cmd.Flags().GetString("email")
			if err != nil {
				printCliErrorAndExit(err)
			}
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				printCliErrorAndExit(err)
			}

			profile := &model.ConfigProfile{
				Name:   args[0],
				APIKey: apiKey,
				Email:  email,
			}

			profileAddService := configServ.NewProfileAddService(configFilepath)
			err = profileAddService.AddNewProfile(profile, force)
			if err != nil {
				printCliErrorAndExit(err)
			}
		},
	}

	cmd.Flags().StringP("api-key", "a", "", "Cloudflare API key")
	cmd.Flags().StringP("email", "e", "", "Cloudflare API email")
	cmd.Flags().StringP("file", "f", "", "The config file to use (If not given the autoloaded config file will be used)")
	cmd.Flags().BoolP("force", "", false, "Force addition of the profile, even if it already exists")

	cmd.MarkFlagRequired("api-key")
	cmd.MarkFlagRequired("email")

	return cmd
}
