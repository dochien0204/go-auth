package handlersLogin

import (
	"fmt"
	"jwt-project/controllers/login"
	"jwt-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service login.Service
}

func NewHandlerLogin(service login.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input login.InputLogin

	ctx.ShouldBindJSON(&input)

	resultLogin, errorLogin := h.service.LoginService(&input)
	fmt.Println("Error", errorLogin)

	if errorLogin == "Invalid account" {
		utils.ValidatorErrorResponse(ctx, "Login Failed", http.StatusBadRequest, http.MethodPost, errorLogin)
		return
	}

	accessTokenData := map[string]interface{}{"username": resultLogin.Username}
	accessToken, errToken := utils.Sign(accessTokenData, "JWT_SECRET", 24*60*1)
	if errToken != nil {
		utils.APIResponse(ctx, "Generate token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, accessToken)
}
