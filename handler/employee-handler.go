package handler

import (
	"crud_api_company/controller/employee"
	"crud_api_company/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerEmployee struct {
	service employee.IEmployee
}

func NewEmployeeHandler(service employee.IEmployee) *handlerEmployee {
	return &handlerEmployee{
		service: service,
	}
}

func (h *handlerEmployee) GetAllEmployee(ctx *gin.Context) {
	result, err := h.service.GetAllEmployee()

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get All Employee Failed", http.StatusBadRequest, http.MethodGet, err)
		return
	}

	utils.APIResponse(ctx, "Get Successfully", http.StatusOK, http.MethodGet, result)
}
