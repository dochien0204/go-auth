package routes

import (
	"jwt-project/controllers/forgot"
	getallusers "jwt-project/controllers/get-all-users"
	"jwt-project/controllers/login"
	"jwt-project/controllers/register"
	forgotHandler "jwt-project/handlers/forgot"
	getallusersHanlder "jwt-project/handlers/get-all-users"
	handlersLogin "jwt-project/handlers/login"
	logoutHandler "jwt-project/handlers/logout"
	registerHandler "jwt-project/handlers/register"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(route *gin.Engine) {

	LoginRepository := login.NewLoginRepository()
	loginService := login.NewServiceLogin(LoginRepository)
	loginHandler := handlersLogin.NewHandlerLogin(loginService)

	registerRepository := register.NewRegisterRepository()
	registerService := register.NewRegisterService(registerRepository)
	registerHandler := registerHandler.NewHandlerRegister(registerService)

	forgotRepository := forgot.NewForgotRepository()
	forgotService := forgot.NewForgotService(forgotRepository)
	forgotHandler := forgotHandler.NewForgotHandler(forgotService)

	getAllUsersRepository := getallusers.NewRepository()
	getAllUsersService := getallusers.NewUserService(getAllUsersRepository)
	getAllUsersHandler := getallusersHanlder.NewGetAllUserHanlder(getAllUsersService)

	groupRoute := route.Group("/api/v1/account")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/logout", logoutHandler.LogoutHandler)
	groupRoute.POST("/signin", loginHandler.LoginHandler)
	groupRoute.POST("/forgot", forgotHandler.ForgotHandler)

	groupRoute.GET("/find-all", getAllUsersHandler.GetAllUserHandler)
	groupRoute.GET("/:username/find", getAllUsersHandler.GetUserByUserName)
}
