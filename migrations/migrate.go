package migrations

import (
	"gorest/conn"
	"gorest/model"
	"log"
)

func UpMigrate() {
	d := conn.DB.AutoMigrate(model.Models...)
	if len(d.GetErrors()) > 0 {
		log.Fatal(d.GetErrors())
	}
}

func DownMigrate() {
	d := conn.DB.DropTableIfExists(model.Models...)
	if len(d.GetErrors()) > 0 {
		log.Fatal(d.GetErrors())
	}
}
