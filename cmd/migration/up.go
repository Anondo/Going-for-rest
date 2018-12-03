package migration

import (
	"going_rest/migrations"
	"log"

	"github.com/spf13/cobra"
)

var UpCommand = cobra.Command{
	Use:   "up",
	Short: "Populates the database",
	Run:   upDatabase,
}

func upDatabase(cmd *cobra.Command, args []string) {
	log.Println("Populating database....")

	migrations.UpMigrate()

	log.Println("Database populated successfully")
}
