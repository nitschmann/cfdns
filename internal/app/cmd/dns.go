package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type dnsCmd struct {
	cmd *cobra.Command
}

func newDNSCmd() *dnsCmd {
	cmd := &cobra.Command{
		Use:     "dns",
		Aliases: []string{"d"},
		Short:   "Cloudflare DNS record management",
		Long:    "Allows to manage DNS records of the Cloudflare account",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	return &dnsCmd{cmd: cmd}
}

func (cmd *dnsCmd) loadSubCommands() {
	cmd.cmd.AddCommand(newDNSAddCmd())
	cmd.cmd.AddCommand(newDNSDeleteCmd())
	cmd.cmd.AddCommand(newDNSUpdateToPublicIPV4Cmd())
	cmd.cmd.AddCommand(newDNSRecordsCmd())
}
