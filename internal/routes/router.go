package routes

import "github.com/gin-gonic/gin"

func (rc *RefRouteController) RefRoutes(router *gin.Engine) {
	auth := router.Group("auth")
	{
		auth.POST("/sign-up", rc.refCont.SignUp) // Регистрация (так же содержит регистрацию с реферальным кодом)
		auth.POST("/sign-in", rc.refCont.SignIn) // Авторизация
	}

	ref := router.Group("referrals")
	{
		ref.POST("/create-ref/:id")   // Создание реферального кода по id
		ref.DELETE("/delete-ref/:id") // Удаление реферального кода по id
		ref.GET("/token")             // Получение реферального кода по email
		ref.GET("/ref-list/:id")      // Получение информации о рефералах по id
	}
}
