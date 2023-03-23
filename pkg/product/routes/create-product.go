package routes

import (
	"Microservices/pkg/product/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {

	body := CreateProductRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Price: body.Price,
		Total: body.Total,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
