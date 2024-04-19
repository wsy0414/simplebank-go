package main

import (
	"database/sql"
	"simplebank/api"
	"simplebank/api/controller"
	"simplebank/api/service"
	"simplebank/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "simplebank/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

//	@title			Swagger SimpleBank API
//	@version		1.0
//	@description	This is a simple bank server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()

	validatorRegister()

	db := connDB()

	implController(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(config.ConfigVal.Server.Port)
}

// connDB return a sql package's DB implement
func connDB() *sql.DB {
	db, err := sql.Open(config.ConfigVal.Database.Driver, config.ConfigVal.Database.Source)
	if err != nil {
		panic(err.Error())
	}

	return db
}

// implController use to dynamic import contorller
func implController(router *gin.Engine, db *sql.DB) {
	// DI
	userService := service.NewUserService(db)
	controller.NewUserController(router, userService)

	activityService := service.NewActivityService(db)
	controller.NewActivityController(activityService, router)
}

// validatorRegister use to register custom validator
func validatorRegister() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordValidate", api.PasswordValidator)
	}
}
