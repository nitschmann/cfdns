package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	configServ "github.com/nitschmann/cfdns/internal/app/service/config"
)

func newConfigCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [PATH]",
		Short: "Create a new config file",
		Long: `
Create a new config file under the given path with the needed structure. If the path is not given, it
will create the file automatically under ~/.cfdns/config (home dir of the current user).
		`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var path string
			if len(args) != 0 {
				path = args[0]
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

			zone, err := cmd.Flags().GetString("zone")
			if err != nil {
				printCliErrorAndExit(err)
			}

			createConfigService := configServ.NewCreateService(apiKey, email, zone)
			configFilepath, err := createConfigService.Create(path, force)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(configFilepath)
		},
	}

	cmd.Flags().StringP("api-key", "a", "", "API Key")
	cmd.MarkFlagRequired("api-key")
	cmd.Flags().StringP("email", "e", "", "API Email")
	cmd.MarkFlagRequired("email")

	cmd.Flags().BoolP("force", "f", false, "Force create (even if file already exists)")
	cmd.Flags().StringP("zone", "z", "", "The default Cloudflare zone to be used")

	return cmd
}
