package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	cloudflareServ "github.com/nitschmann/cfdns/internal/app/service/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/util/cmdhelper"
)

func newDnsAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [ZONE_ID_OR_NAME]",
		Aliases: []string{"a", "add-record"},
		Short:   "Add a new DNS record to a Cloudflare zone",
		Long: `
Add a new DNS record to a Cloudflare zone with the specified parameters. The zone could be either identified by its ID or name.
If successful, it prints the ID, type and name of the newly created DNS record. [Format: __ID__,__type__,__name__]`,
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

			zone, err := zoneService.FindByIdOrName(args[0])
			if err != nil {
				printCliErrorAndExit(err)
			}

			dnsService, err := cloudflareServ.NewDnsService(cloudflareConfig)
			if err != nil {
				printCliErrorAndExit(err)
			}

			// Attributes
			zoneID := zone.ID
			dnsRecordType, err := cmd.Flags().GetString("type")
			if err != nil {
				printCliErrorAndExit(err)
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				printCliErrorAndExit(err)
			}
			content, err := cmd.Flags().GetString("content")
			if err != nil {
				printCliErrorAndExit(err)
			}
			ttl, err := cmd.Flags().GetInt("ttl")
			if err != nil {
				printCliErrorAndExit(err)
			}
			priority, err := cmd.Flags().GetInt("priority")
			if err != nil {
				printCliErrorAndExit(err)
			}
			proxied, err := cmd.Flags().GetBool("proxied")
			if err != nil {
				printCliErrorAndExit(err)
			}

			dnsRecord, err := dnsService.Create(zoneID, dnsRecordType, name, content, ttl, priority, proxied)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(strings.Join([]string{dnsRecord.ID, dnsRecord.Type, dnsRecord.Name}, ","))
		},
	}

	// auth flags
	cmd.Flags().StringP("api-key", "a", "", "Cloudflare API key")
	cmd.Flags().StringP("email", "e", "", "Cloudflare email")
	cmd.Flags().StringP("profile", "p", "default", "Config profile to use")
	// entity flags
	cmd.Flags().StringP("type", "", "", "DNS record type")
	cmd.MarkFlagRequired("type")
	cmd.Flags().StringP("name", "", "", "DNS record name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringP("content", "", "", "DNS record content")
	cmd.MarkFlagRequired("content")
	cmd.Flags().IntP("ttl", "", 1, "Time to live for DNS record")
	cmd.Flags().IntP("priority", "", 0, "Used with some records like MX and SRV to determine priority")
	cmd.Flags().BoolP("proxied", "", false, "Whether the record is receiving the performance and security benefits of Cloudflare")

	return cmd
}
