package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	publicIpServ "github.com/nitschmann/cfdns/internal/app/service/publicip"
	"github.com/nitschmann/cfdns/pkg/util/httpclient"
)

func newPublicIPV4Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "public-ipv4",
		Aliases: []string{"ipv4", "IPv4", "outgoing-ip", "public-ip"},
		Short:   "Print the public outgoing IPv4 address",
		Long:    "Fetches (through network) the public and outgoing IPv4 address of the machine",
		Run: func(cmd *cobra.Command, args []string) {
			httpClient := httpclient.New()
			publicIPService := publicIpServ.New(httpClient)

			publicIPV4, err := publicIPService.FetchPublicIPV4()
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(publicIPV4)
		},
	}

	return cmd
}
