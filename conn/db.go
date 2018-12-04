package conn

import (
	"going_rest/checker"
	"going_rest/config"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	var err error

	d := config.GetDB()
	name := d.Name
	username := d.Username
	protocol := d.Protocol
	host := d.Host
	port := d.Port
	db := d.DB

	connectionString := username + ":@" + protocol + "(" + host + ":" + strconv.Itoa(port) +
		")/" + db + "?parseTime=true"

	DB, err = gorm.Open(name, connectionString)
	checker.CheckErr(err)

}
