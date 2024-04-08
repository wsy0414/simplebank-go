package main

import (
	"log"
	"simplebank/config"
	"simplebank/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.HelloWorld)

	c, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	router.Run(c.Server.Port)
}
