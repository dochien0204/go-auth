package userhandler

import (
	"crud-api/controllers"
	"crud-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service controllers.IUserService
}

func NewUserHandler(service controllers.IUserService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateNewUserHanlder(ctx *gin.Context) {
	var inputCreate controllers.UserCreateInput

	ctx.ShouldBindJSON(&inputCreate)

	resultCreate, errCreate := h.service.CreateNewUser(&inputCreate)

	if errCreate != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Created Failed", http.StatusBadRequest, http.MethodPost, errCreate)
		return
	}

	utils.APIResponse(ctx, "Created Successfully", http.StatusOK, http.MethodPost, resultCreate)
}
