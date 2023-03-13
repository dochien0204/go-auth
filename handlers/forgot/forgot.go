package forgotHandler

import (
	"jwt-project/controllers/forgot"
	"jwt-project/presenters"
	"jwt-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service forgot.IForgotService
}

func NewForgotHandler(service forgot.IForgotService) *handler {
	return &handler{service: service}
}

func (h *handler) ForgotHandler(ctx *gin.Context) {
	var input forgot.InputForgot
	ctx.ShouldBindJSON(&input)

	resultForgot, errForgot := h.service.ForgotService(&input)

	if !errForgot {
		utils.ValidatorErrorResponse(ctx, "Failed", http.StatusBadRequest, http.MethodPost, resultForgot)
		return
	}

	forgotPasswordResponse := presenters.ForgotPasswordResponse{Result: resultForgot}
	utils.APIResponse(ctx, "Successfully", http.StatusOK, http.MethodPost, forgotPasswordResponse)
}
