package migration

import (
	"gorest/migrations"
	"log"

	"github.com/spf13/cobra"
)

var (
	DownCommand = &cobra.Command{
		Use:   "down",
		Short: "Depopulates the database",
		Run:   downDatabase,
	}
)

func downDatabase(cmd *cobra.Command, args []string) {
	log.Println("Depopulating database....")

	migrations.DownMigrate()

	log.Println("Database depopulated successfully")
}
