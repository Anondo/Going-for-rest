package migration

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmnd := cobra.Command{
		Use:   "migration",
		Short: "Run Database Migrations",
	}

	return &rootCmnd
}

func init() {
	r := RootCommand()
	r.AddCommand(&UpCommand)
	r.AddCommand(&DownCommand)
}
