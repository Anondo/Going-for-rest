package cmd

import (
	"going_rest/cmd/migration"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmnd := cobra.Command{
		Use:   "gorest",
		Short: "gorest is a cli app for a rest application",
	}

	return &rootCmnd
}

func init() {
	r := RootCommand()
	r.AddCommand(migration.RootCommand())
}
