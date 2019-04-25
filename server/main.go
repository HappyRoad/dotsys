package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"os"
)

func HeartBeat(c *gin.Context){
	var data dotsys.Request
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}


var (
	port = flag.Int("port", 7777, "port")
	help = flag.Bool("h", false, "help")
	mysql = flag.String("mysql", "127.0.0.1:3306", "mysql address")
)

func main() {
	flag.Parse()
	if len(os.Args) == 1 || *help {
		flag.Usage()
		os.Exit(0)
	}
	r := gin.Default()

	r.POST("/heartbeat", HeartBeat)

	r.Run(fmt.Sprintf(":%d", *port))
}
