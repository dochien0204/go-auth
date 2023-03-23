package product

import (
	"Microservices/pkg/config"
	"Microservices/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	r.Group("/product")
	r.POST("/", svc.CreateProduct)

	return svc
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
