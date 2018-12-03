package migration

import (
	"github.com/spf13/cobra"
)

var (
	RootCommand = &cobra.Command{
		Use:   "migration",
		Short: "Run Database Migrations",
	}
)

func init() {
	RootCommand.AddCommand(UpCommand)
	RootCommand.AddCommand(DownCommand)
}
