package cli

import "github.com/spf13/cobra"

type ServeFlags struct {
	Web struct {
		Addr string
	}
	Store StoreConfig
}

func cmdServe() *cobra.Command {
	var flags ServeFlags
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the golinks server",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.Flags().StringVar(&flags.Web.Addr, "web.addr", ":8080", "HTTP server address")
	cmd.Flags().Var(&flags.Store, "store", "Store type (memory, sql, leveldb)")

	return cmd
}
