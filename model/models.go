package model

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	Title string `gorm:"unique;not null"`
	Body  string
}
