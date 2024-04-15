package main

import (
	"log"
	"simplebank/config"
	"simplebank/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	loadRouter(router)

	c, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	router.Run(c.Server.Port)
}

func loadRouter(router *gin.Engine) {
	router.POST("/signup", controller.SignUp)

}
