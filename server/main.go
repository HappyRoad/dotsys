package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var (
	port  = flag.Int("port", 7777, "port")
	help  = flag.Bool("h", false, "help")
	mysql = flag.String("mysql", "root:root@(127.0.0.1:3306)", "mysql address")
)

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	r := gin.Default()

	server, err := NewServer(*mysql)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.POST("/heartbeat", server.HeartBeat)

	r.Run(fmt.Sprintf(":%d", *port))
}
