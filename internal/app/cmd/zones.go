package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	cloudflareServ "github.com/nitschmann/cfdns/internal/app/service/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/util/cmdhelper"
)

func newZonesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "zones",
		Aliases: []string{"z"},
		Short:   "List all Cloudflare zones for the configuration",
		Long: `
Prints all the Cloudflare zones with their most important details in a table.
It uses the configuration of the specified profile or directly the api-key and email flag (if set).
		`,
		Run: func(cmd *cobra.Command, args []string) {
			cloudflareConfig, err := cmdhelper.GetCloudflareConfigByFlags(cmd)
			if err != nil {
				printCliErrorAndExit(err)
			}

			zoneService, err := cloudflareServ.NewZoneService(cloudflareConfig)
			if err != nil {
				printCliErrorAndExit(err)
			}

			zones, err := zoneService.List()
			if err != nil {
				printCliErrorAndExit(err)
			}

			withoutTable, err := cmd.Flags().GetBool("without-table")
			if err != nil {
				printCliErrorAndExit(err)
			}

			columnNames := []string{"ID", "Name", "Type", "Status", "Created At", "Modified On"}
			table := cmdhelper.TableWithHeader(columnNames)

			if withoutTable {
				fmt.Println(strings.Join(columnNames, ","))
			}

			for _, z := range zones {
				line := []string{
					z.ID,
					z.Name,
					z.Type,
					z.Status,
					z.CreatedOn.String(),
					z.ModifiedOn.String(),
				}

				if withoutTable {
					fmt.Println(strings.Join(line, ","))
				} else {
					table.Append(line)
				}

			}

			if !withoutTable {
				table.Render()
			}

		},
	}

	cmd.Flags().StringP("api-key", "a", "", "Cloudflare API key")
	cmd.Flags().StringP("email", "e", "", "Cloudflare email")
	cmd.Flags().StringP("profile", "p", "default", "Config profile to use")

	cmd.Flags().BoolP("without-table", "", false, "Print list without table")

	return cmd
}
