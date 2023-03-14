package userhandler

import (
	"crud-api/controllers"
	"crud-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateUserByIdHandler(ctx *gin.Context) {

	param := ctx.Param("userId")
	paramConv, errConv := strconv.Atoi(param)

	if errConv != nil {
		utils.ValidatorErrorResponse(ctx, "Update Failed", http.StatusBadRequest, http.MethodPatch, "Id is invalid")
		return
	}
	var input controllers.UserUpdateInput

	ctx.ShouldBindJSON(&input)

	result, err := h.service.UpdateUserById(uint(paramConv), &input)

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Update Failed", http.StatusBadRequest, http.MethodPatch, err)
		return
	}

	utils.APIResponse(ctx, "Update Successfully", http.StatusOK, http.MethodPatch, result)
}
