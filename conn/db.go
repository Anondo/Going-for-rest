package conn

import (
	"fmt"
	"going_rest/checker"
	"going_rest/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect() {
	var err error

	d := config.GetDB()
	name := d.Name
	username := d.Username
	password := d.Password
	host := d.Host
	port := d.Port
	db := d.DB

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, username, db, password)

	DB, err = gorm.Open(name, connectionString)
	checker.CheckErr(err)

}
