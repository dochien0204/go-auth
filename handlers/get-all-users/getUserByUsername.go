package getallusersHanlder

import (
	"jwt-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetUserByUserName(ctx *gin.Context) {

	userNameParam := ctx.Param("username")
	result, err := h.service.GetUserByUserName(userNameParam)

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get User Failed", http.StatusNotFound, http.MethodGet, err)
		return
	}

	utils.APIResponse(ctx, "Get User Successfully", http.StatusOK, http.MethodGet, result)
}
