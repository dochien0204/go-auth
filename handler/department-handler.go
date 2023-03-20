package handler

import (
	"crud_api_company/controller/department"
	"crud_api_company/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerDepartment struct {
	service department.IDepartment
}

func NewDepartmentHandler(service department.IDepartment) *handlerDepartment {
	return &handlerDepartment{
		service: service,
	}
}

func (h *handlerDepartment) GetAllDepartMentHandler(ctx *gin.Context) {

	result, err := h.service.GetAllDepartMent()

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get All Department Failed", http.StatusBadRequest, http.MethodGet, err)
		return
	}

	utils.APIResponse(ctx, "Get all department successfully", http.StatusOK, http.MethodGet, result)
}
