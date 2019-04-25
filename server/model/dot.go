package model

import (
	"github.com/jinzhu/gorm"
)

type Dot struct {
	gorm.Model

	Project string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`
	Host string `gorm:"unique;not null"`
	Ip string

}
