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
	userGroupRoute.GET("list", userHandler.GetAllUserHandler)
	userGroupRoute.GET("/:userId/find", userHandler.GetUserByIdHandler)
	userGroupRoute.PATCH("/:userId/update", userHandler.UpdateUserByIdHandler)
	userGroupRoute.DELETE("/:userId/delete", userHandler.DeleteUserByIdHandler)
}
