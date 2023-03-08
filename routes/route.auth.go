package routes

import (
	"jwt-project/controllers/login"
	handlersLogin "jwt-project/handlers/login"
	logoutHandler "jwt-project/handlers/logout"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(route *gin.Engine) {

	LoginRepository := login.NewLoginRepository()
	loginService := login.NewServiceLogin(LoginRepository)
	loginHandler := handlersLogin.NewHandlerLogin(loginService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/logout", logoutHandler.LogoutHandler)
	groupRoute.POST("/signin", loginHandler.LoginHandler)
}
