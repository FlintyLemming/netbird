package cmd

import (
	"github.com/spf13/cobra"

	"github.com/FlintyLemming/netbird/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "prints Netbird version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.Println(version.NetbirdVersion())
		},
	}
)
