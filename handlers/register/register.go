package registerHandler

import (
	"jwt-project/controllers/register"
	"jwt-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service register.IService
}

func NewHandlerRegister(service register.IService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var input register.InputRegister

	ctx.ShouldBindJSON(&input)

	resultRegister, errRegister := h.service.RegisterService(&input)

	if errRegister == "Account already exists" {
		utils.ValidatorErrorResponse(ctx, "Register failed", http.StatusBadRequest, http.MethodPost, errRegister)
		return
	}

	utils.APIResponse(ctx, "Registersuccessfully", http.StatusOK, http.MethodPost, resultRegister)
}
