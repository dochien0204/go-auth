package logoutHandler

import (
	"jwt-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(ctx *gin.Context) {
	// token := ctx.Param("token")

	resultToken, errToken := utils.VerifyTokenHeader(ctx, "JWT_SECRET")

	if errToken != nil {
		utils.APIResponse(ctx, "Verified token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := utils.DecodeToken(resultToken)

	result.Claims.Active = false
	result.Claims.Authorization = false

	utils.APIResponse(ctx, "Logout succesfully", http.StatusOK, http.MethodPost, result)
}
