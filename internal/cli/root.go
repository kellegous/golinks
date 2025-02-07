package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use: "golinks",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}

func Execute() {
	if err := cmdRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
