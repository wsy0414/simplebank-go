package main

import (
	"simplebank/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.HelloWorld)

	router.Run(":8080")
}
