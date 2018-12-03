package conn

import (
	"going_rest/checker"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?parseTime=true")
	checker.CheckErr(err)
}
