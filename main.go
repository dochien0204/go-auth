package main

import (
	"crud_api_company/config"
	"crud_api_company/model"
	"crud_api_company/route"

	"github.com/gin-gonic/gin"
)

func main() {
	//connect DB
	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.Company{}, &model.Department{})

	//definite route
	router := gin.Default()

	route.InItCompanyRoute(router)
	router.Run(":8080")

}
