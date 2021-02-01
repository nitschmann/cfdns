package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	cloudflareServ "github.com/nitschmann/cfdns/internal/app/service/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/util/cmdhelper"
)

func newDNSDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [ZONE_ID_OR_NAME] [DNS_RECORD_ID_OR_NAME]",
		Aliases: []string{"d", "delete-record"},
		Short:   "Delete a DNS record from a Cloudflare zone",
		Long: `
Delete a DNS record from a Cloudflare zone. The zone could be either identified by its ID or name.
The DNS record is identified by its ID or name. But attention: There could be potentially multiple DNS records for the same name, so a type flag should be passed as well.
The operation will fail if there are still multiple recrods available. It is therefore recommended to work with the ID of the record.
If successful, it prints the ID, type and name of the deleted DNS record. [Format: __ID__,__type__,__name__]`,
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

			zone, err := zoneService.FindByIDOrName(args[0])
			if err != nil {
				printCliErrorAndExit(err)
			}

			dnsService, err := cloudflareServ.NewDNSService(cloudflareConfig)
			if err != nil {
				printCliErrorAndExit(err)
			}

			// Attributes
			dnsRecordID := args[1]
			dnsRecordType, err := cmd.Flags().GetString("type")
			if err != nil {
				printCliErrorAndExit(err)
			}

			deletedDNSRecord, err := dnsService.DeleteByIDOrNameAndType(zone, dnsRecordID, dnsRecordType)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(strings.Join([]string{deletedDNSRecord.ID, deletedDNSRecord.Type, deletedDNSRecord.Name}, ","))

		},
	}

	cmd.Flags().StringP("api-key", "a", "", "Cloudflare API key")
	cmd.Flags().StringP("email", "e", "", "Cloudflare email")
	cmd.Flags().StringP("profile", "p", "default", "Config profile to use")

	cmd.Flags().StringP("type", "", "", "DNS record type")

	return cmd
}
