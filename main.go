package main

import (
	"jwt-project/configs"
	"jwt-project/models"
	"jwt-project/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.ConnectDB()
	configs.Database.AutoMigrate(&models.EntityUser{})
	router := gin.Default()

	routes.InitAuthRoutes(router)
	router.Run("localhost:8080")
	log.Fatal()

}
