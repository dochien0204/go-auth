package main

import (
	"crud-api/configs"
	"crud-api/models"
	"crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	configs.DB.AutoMigrate(&models.Users{})

	router := gin.Default()
	routes.InitUseRoutes(router)
	router.Run("localhost:8080")
}
