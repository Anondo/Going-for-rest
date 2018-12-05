package cmd

import (
	"gorest/server"

	"github.com/spf13/cobra"
)

var (
	ServeCommand = &cobra.Command{
		Use:   "serve",
		Short: "Starts the application server",
		Run:   startServer,
	}
)

func startServer(cmd *cobra.Command, args []string) {
	server.Start()
}
