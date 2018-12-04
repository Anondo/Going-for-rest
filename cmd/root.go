package cmd

import (
	"going_rest/checker"
	"going_rest/cmd/migration"
	"going_rest/config"
	"going_rest/conn"

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
	RootCommand.Flags().IntP("port", "p", 0, "the port to do things on")
	RootCommand.Flags().StringP("config", "c", "", "the configuration file")

	cobra.OnInitialize(initConfig)

	RootCommand.AddCommand(migration.RootCommand)
	RootCommand.AddCommand(ServeCommand)
}

func initConfig() {
	config.LoadConfig(RootCommand)
	conn.Connect()
}
