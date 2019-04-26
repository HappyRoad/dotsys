package main

import (
	"fmt"
	"github.com/HappyRoad/dotsys"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"server/db"
	"server/model"
)

type Server struct {
	db        *gorm.DB
	mdataDB   *db.PrometheusDB
}

func NewServer(mysql string) (server *Server, err error) {
	server = &Server{}
	server.db, err = gorm.Open("mysql", fmt.Sprintf("%s/dotsys?charset=utf8&parseTime=True&loc=Local", mysql))
	if err != nil {
		return
	}
	server.mdataDB, err = db.NewPrometheusDB()
	if err != nil {
		return
	}
	server.db = server.db.Table("dot").Debug()
	return
}

func (server *Server) HeartBeat(c *gin.Context) {
	var (
		data dotsys.Request
		err error
	)
	if err = c.ShouldBindJSON(&data); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	dot := model.NewDot(data)
	err = server.db.Model(dot).Select("updated_at").Update(dot).Error
	if err != nil {
		err = server.db.Create(dot).Error
	}
	if err != nil {
		log.Printf("update db error: %s", err.Error())
	}
	// 存储打点数据
	if len(data.DotValues) > 0 {
		server.mdataDB.Save(&data)
	}
}
