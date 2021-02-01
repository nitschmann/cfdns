package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	cloudflareServ "github.com/nitschmann/cfdns/internal/app/service/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/util/cmdhelper"
)

func newDNSRecordsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "records [ZONE_ID_OR_NAME]",
		Aliases: []string{"list", "r"},
		Short:   "Print list of all DNS records for a Cloudflare zone",
		Long: `
Print list of all DNS records for a Cloudflare zone in a table. The zone could be either identified by its ID or name.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cloudflareConfig, err := cmdhelper.GetCloudflareConfigByFlags(cmd)
			if err != nil {
				printCliErrorAndExit(err)
			}

			zoneService, err := cloudflareServ.NewZoneService(cloudflareConfig)
			if err != nil {
				printCliErrorAndExit(err)
			}

			zone, err := zoneService.FindByIDOrName(args[0])
			if err != nil {
				printCliErrorAndExit(err)
			}

			dnsService, err := cloudflareServ.NewDNSService(cloudflareConfig)
			if err != nil {
				printCliErrorAndExit(err)
			}

			dnsRecords, err := dnsService.List(zone.ID)
			if err != nil {
				printCliErrorAndExit(err)
			}

			withoutTable, err := cmd.Flags().GetBool("without-table")
			if err != nil {
				printCliErrorAndExit(err)
			}

			columnNames := []string{
				"ID",
				"Type",
				"Name",
				"Content",
				"TTL",
				"Priority",
				"Proxied",
				"Created On",
				"Modified On",
			}
			table := cmdhelper.TableWithHeader(columnNames)

			if withoutTable {
				fmt.Println(strings.Join(columnNames, ","))
			}

			for _, d := range dnsRecords {
				line := []string{
					d.ID,
					d.Type,
					d.Name,
					d.Content,
					strconv.Itoa(d.TTL),
					strconv.Itoa(d.Priority),
					strconv.FormatBool(d.Proxied),
					d.CreatedOn.String(),
					d.ModifiedOn.String(),
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
