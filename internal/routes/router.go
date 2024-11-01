package routes

import (
	_ "ReferralAPI/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (rc *RefRouteController) RefRoutes(router *gin.Engine) {
	auth := router.Group("auth")
	{
		auth.POST("/sign-up", rc.refCont.SignUp) // Регистрация (так же содержит регистрацию с реферальным кодом)
		auth.POST("/sign-in", rc.refCont.SignIn) // Авторизация
	}

	ref := router.Group("referrals")
	{
		ref.POST("/create-ref", rc.refCont.CreateRef)   // Создание реферального кода
		ref.DELETE("/delete-ref", rc.refCont.DeleteRef) // Удаление реферального кода
		ref.GET("/token", rc.refCont.GetToken)          // Получение реферального кода по email
		ref.GET("/ref-list/:id", rc.refCont.GetRefList) // Получение информации о рефералах по id
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
