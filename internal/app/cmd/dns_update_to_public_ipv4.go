package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	cloudflareServ "github.com/nitschmann/cfdns/internal/app/service/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/util/cmdhelper"
)

func newDnsUpdateToPublicIpV4Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-to-public-ipv4 [ZONE_ID_OR_NAME] [DNS_RECORD_ID_OR_NAME]",
		Aliases: []string{"IPv4"},
		Short:   "Update a Cloudflare DNS A record content to the public IPv4 of this machine",
		Long: `
Update a Cloudflare DNS A record content to the public IPv4 of this machine. The zone could be either identified by its ID or name. The DNS record is identified by its ID or name. The operation will fail if there are still multiple recrods available or if the DNS record is not a type A record. It is therefore recommended to work with the ID of the record.
If successful, it prints the ID, type and name of the updated DNS record. [Format: __ID__,__type__,__name__]`,
		Args: cobra.ExactArgs(2),
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

			dnsRecordID := args[1]
			dnsRecord, err := dnsService.UpdateARecordContentToPublicIpV4(zone, dnsRecordID)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(strings.Join([]string{dnsRecord.ID, dnsRecord.Type, dnsRecord.Name}, ","))

		},
	}

	cmd.Flags().StringP("api-key", "a", "", "Cloudflare API key")
	cmd.Flags().StringP("email", "e", "", "Cloudflare email")
	cmd.Flags().StringP("profile", "p", "default", "Config profile to use")

	return cmd
}
