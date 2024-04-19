package main

import (
	"database/sql"
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

	validatorRegister()

	db := connDB()

	implController(router, db)

	router.Run(config.ConfigVal.Server.Port)
}

func connDB() *sql.DB {
	db, err := sql.Open(config.ConfigVal.Database.Driver, config.ConfigVal.Database.Source)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func implController(router *gin.Engine, db *sql.DB) {
	// DI
	userService := service.NewUserService(db)
	controller.NewUserController(router, userService)

	activityService := service.NewActivityService(db)
	controller.NewActivityController(activityService, router)
}

func validatorRegister() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordValidate", api.PasswordValidator)
	}
}
