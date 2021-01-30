package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nitschmann/cfdns/internal/pkg/config"
)

var (
	// AppVersion is the global CLI application version
	AppVersion string
	rootCmd    *RootCmd

	_ = func() error {
		rootCmd = NewRootCmd()
		rootCmd.LoadSubCommands()

		err := config.SetUpLoader()
		if err != nil {
			printCliErrorAndExit(err)
		}

		return nil
	}()
)

// RootCmd is a global cmd package abstraction struct
type RootCmd struct {
	Cmd *cobra.Command
}

// Execute is the app-wide CLI entrypoint
func Execute() {
	err := rootCmd.Cmd.Execute()
	if err != nil {
		printCliErrorAndExit(err)
	}
}

// LoadSubCommands loads the sub-commands of RootCmd.Cmd
func (r *RootCmd) LoadSubCommands() {
	cmd := r.Cmd
	cmd.AddCommand(newPublicIpV4Cmd())
	cmd.AddCommand(newVersionCmd())

	// config command
	configCmd := newConfigCmd()
	configCmd.loadSubCommands()
	cmd.AddCommand(configCmd.cmd)
}

// NewRootCmd returns the application and global facing root cobra command
func NewRootCmd() *RootCmd {
	cmd := &cobra.Command{
		Use:   "cfdns",
		Short: "CLI tool to manage Cloudflare DNS entries",
		Long: `
cfdns is a tool which allows to manage Cloudflare DNS entries via the API easily within the CLI. It also has the option to set dynamically the public IPv4 of the machine (or the network itself), through detection, for DNS entries. A sytemwide config file allows furthermore to work with different profiles (API keys) at the same time. This tool does NOT cover anything else of the Cloudflare API.
		`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := config.Load()
			if err != nil {
				printCliErrorAndExit(err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	return &RootCmd{Cmd: cmd}
}

func printCliErrorAndExit(msg interface{}) {
	fmt.Printf("An unexpected error occurred:\n%s\n", msg)
	os.Exit(1)
}
