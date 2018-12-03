package migrations

import (
	"going_rest/conn"
	"going_rest/model"
)

func UpMigrate() {
	conn.DB.AutoMigrate(&model.Blog{})
}

func DownMigrate() {
	conn.DB.DropTableIfExists(&model.Blog{})
}
