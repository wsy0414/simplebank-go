package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	"simplebank/api/controller"
	"simplebank/api/service"
	"simplebank/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	c, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordValidate", api.PasswordValidator)
	}

	db, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("init")
	// DI
	userService := service.NewUserService(db)
	controller.NewUserController(router, userService)

	router.Run(c.Server.Port)
}
