package userhandler

import (
	"crud-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteUserByIdHandler(ctx *gin.Context) {

	param := ctx.Param("userId")
	paramConv, errConv := strconv.Atoi(param)

	if errConv != nil {
		utils.ValidatorErrorResponse(ctx, "Update Failed", http.StatusBadRequest, http.MethodDelete, "Id is invalid")
		return
	}

	result, err := h.service.DeleteUserById(uint(paramConv))

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Delete Failed", http.StatusBadRequest, http.MethodDelete, err)
		return
	}

	utils.APIResponse(ctx, "Delete successfully", http.StatusOK, http.MethodDelete, result)
}
