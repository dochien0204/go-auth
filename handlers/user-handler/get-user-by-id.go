package userhandler

import (
	"crud-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetUserByIdHandler(ctx *gin.Context) {

	param := ctx.Param("userId")

	convertParam, err := strconv.Atoi(param)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusBadRequest, http.MethodGet, "ID is invalid")
		return
	}

	resultGet, errGet := h.service.GetUserById(uint(convertParam))

	if errGet != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusNotFound, http.MethodGet, errGet)
		return
	}

	utils.APIResponse(ctx, "Get Successfully", http.StatusOK, http.MethodGet, resultGet)
}
