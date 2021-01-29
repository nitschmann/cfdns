package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	publicIpServ "github.com/nitschmann/cfd/internal/app/service/publicip"
	"github.com/nitschmann/cfd/pkg/util/httpclient"
)

func newPublicIpV4Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "public-ipv4",
		Aliases: []string{"ipv4", "IPv4", "outgoing-ip", "public-ip"},
		Short:   "Print the public outgoing IPv4 address",
		Long:    "Fetches (through network) the public and outgoing IPv4 address of the machine",
		Run: func(cmd *cobra.Command, args []string) {
			httpClient := httpclient.New()
			publicIpService := publicIpServ.New(httpClient)

			publicIpV4, err := publicIpService.FetchPublicIpV4()
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(publicIpV4)
		},
	}

	return cmd
}
