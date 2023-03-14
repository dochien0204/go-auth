package getallusersHanlder

import (
	getallusers "jwt-project/controllers/get-all-users"
	"jwt-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateUserByIdHandler(ctx *gin.Context) {

	var input getallusers.InputUpdateUser

	ctx.ShouldBindJSON(&input)

	param := ctx.Param("userId")

	userIdParam, _ := strconv.Atoi(param)

	resultUpdate, errUpdate := h.service.UpdateUserById(uint(userIdParam), &input)

	if errUpdate != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Updated failed", http.StatusBadRequest, http.MethodPatch, errUpdate)
		return
	}

	utils.APIResponse(ctx, "Updated successfully", http.StatusOK, http.MethodPatch, resultUpdate)
}
