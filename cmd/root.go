package cmd

import (
	"going_rest/checker"
	"going_rest/cmd/migration"

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
	RootCommand.AddCommand(migration.RootCommand)
	RootCommand.AddCommand(ServeCommand)
}
