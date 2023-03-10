package main

import (
	"jwt-project/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.InitAuthRoutes(router)
	router.Run("localhost:8080")
	log.Fatal()

}
