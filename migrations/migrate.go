package migrations

import (
	"going_rest/conn"
	"going_rest/model"
	"log"
)

func UpMigrate() {
	d := conn.DB.AutoMigrate(&model.Blog{})
	if len(d.GetErrors()) > 0 {
		log.Fatal(d.GetErrors())
	}
}

func DownMigrate() {
	d := conn.DB.DropTableIfExists(&model.Blog{})
	if len(d.GetErrors()) > 0 {
		log.Fatal(d.GetErrors())
	}
}
