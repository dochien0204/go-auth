package getallusersHanlder

import (
	"jwt-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteUserByIdHandler(ctx *gin.Context) {

	var param = ctx.Param("userId")

	convertParam, errConvert := strconv.Atoi(param)

	if errConvert != nil {
		utils.ValidatorErrorResponse(ctx, "Something went wrong", http.StatusBadRequest, http.MethodDelete, errConvert)
		return
	}

	resultDelete, errDelete := h.service.DeleteUserById(uint(convertParam))

	if errDelete != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Deleted Failed", http.StatusBadRequest, http.MethodDelete, errDelete)
		return
	}

	utils.APIResponse(ctx, "Deleted successfully", http.StatusOK, http.MethodDelete, resultDelete)
}
