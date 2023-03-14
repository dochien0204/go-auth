package userhandler

import (
	"crud-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllUserHandler(ctx *gin.Context) {

	listUser, errGetAll := h.service.GetAllUser()

	if errGetAll != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get All User Failed", http.StatusBadRequest, http.MethodGet, errGetAll)
		return
	}

	utils.APIResponse(ctx, "Get All User Successfully", http.StatusOK, http.MethodGet, listUser)
}
