package main

import (
	"crud-api/configs"
	"crud-api/models"
	"crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// config := consulapi.DefaultConfig()

	// consul, _ := consulapi.NewClient(config)

	// services, _ := consul.Agent().Services()

	// companyService := services["company_api"]

	// address := companyService.Address
	// port := companyService.Port
	// url := fmt.Sprintf("http://%s:%v", address, port)
	// fmt.Println("Link:", url)
	configs.ConnectDB()
	configs.DB.AutoMigrate(&models.Users{})

	router := gin.Default()
	routes.InitUseRoutes(router)
	router.Run(":8081")

}
