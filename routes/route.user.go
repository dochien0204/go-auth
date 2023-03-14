package routes

import (
	"crud-api/controllers"
	userhandler "crud-api/handlers/user-handler"

	"github.com/gin-gonic/gin"
)

func InitUseRoutes(route *gin.Engine) {

	userGroupRoute := route.Group("/api/v1/users")

	userRepository := controllers.NewRepository()
	userService := controllers.NewUserService(userRepository)
	userHandler := userhandler.NewUserHandler(userService)

	// userGroupRoute.GET("/list", userHandler.CreateNewUserHanlder)
	userGroupRoute.POST("/create-user", userHandler.CreateNewUserHanlder)
}
