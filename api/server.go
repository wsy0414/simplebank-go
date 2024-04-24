package api

import (
	"simplebank/api/controller"
	"simplebank/api/service"
	"simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(store sqlc.Store) *gin.Engine {
	router := gin.Default()

	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validatorRegister(trans)

	implController(router, store, trans)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// implController use to dynamic import contorller
func implController(router *gin.Engine, store sqlc.Store, trans ut.Translator) {
	// DI
	userService := service.NewUserService(store)
	controller.NewUserController(router, trans, userService)

	balanceService := service.NewBalanceService(store)
	controller.NewBalanceController(router, balanceService)
}

// validatorRegister use to register custom validator
func validatorRegister(trans ut.Translator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordValidate", PasswordValidator)
		v.RegisterValidation("currencyValidate", CurrencyValidator)

		en_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())

			return t
		})
	}
}
