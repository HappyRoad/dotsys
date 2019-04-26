package model

import (
	"github.com/HappyRoad/dotsys"
	"time"
)

type Dot struct {
	Project string `gorm:"primary_key"`
	Host string `gorm:"primary_key"`
	Ip string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func NewDot(request dotsys.Request)(dot *Dot){
	dot = &Dot{
		Project: request.Project,
		Host: request.Host,
		Ip: request.Ip,
	}

	return
}
