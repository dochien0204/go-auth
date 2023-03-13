package getallusersHanlder

import (
	getallusers "jwt-project/controllers/get-all-users"
	"jwt-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getallusers.Service
}

func NewGetAllUserHanlder(service getallusers.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAllUserHandler(ctx *gin.Context) {

	resultListAccount, err := h.service.GetAllUsers()

	if !err {
		utils.ValidatorErrorResponse(ctx, "Get List Account failed", http.StatusBadRequest, http.MethodGet, "Get list Account failed")
		return
	}

	utils.APIResponse(ctx, "Successfully", http.StatusOK, http.MethodGet, resultListAccount)
}
