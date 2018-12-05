package cmd

import (
	"gorest/checker"
	"gorest/cmd/migration"
	"gorest/config"
	"gorest/conn"

	"github.com/spf13/cobra"
)

var (
	RootCommand = &cobra.Command{
		Use:   "gorest",
		Short: "gorest is a cli app for a rest application",
	}
)

func Execute() {
	checker.CheckErr(RootCommand.Execute())
}

func init() {
	RootCommand.PersistentFlags().IntP("port", "p", 0, "the port to run http server on")
	RootCommand.Flags().StringP("config", "c", "", "the configuration file")

	cobra.OnInitialize(initConfig)

	RootCommand.AddCommand(migration.RootCommand)
	RootCommand.AddCommand(ServeCommand)
}

func initConfig() {
	config.LoadConfig(RootCommand)
	conn.Connect()
}
