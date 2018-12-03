package migration

import (
	"going_rest/migrations"
	"log"

	"github.com/spf13/cobra"
)

var DownCommand = cobra.Command{
	Use:   "up",
	Short: "Depopulates the database",
	Run:   downDatabase,
}

func downDatabase(cmd *cobra.Command, args []string) {
	log.Println("Populating database....")

	migrations.DownMigrate()

	log.Println("Database depopulated successfully")
}
