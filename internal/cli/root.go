package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func cmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use: "golinks",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(cmdServe())

	return cmd
}

func Execute() {
	if err := cmdRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
