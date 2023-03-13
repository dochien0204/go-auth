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

	if errForgot != "Forgot successfully" {
		utils.ValidatorErrorResponse(ctx, "Forgot failed", http.StatusBadRequest, http.MethodPost, errForgot)
		return
	}

	resetLink := "http://localhost:8080/" + resultForgot.Username + "/reset-password"

	forgotPasswordResponse := presenters.ForgotPasswordResponse{PasswordResetLink: resetLink}
	utils.APIResponse(ctx, "Forgot successfully", http.StatusOK, http.MethodPost, forgotPasswordResponse)
}
