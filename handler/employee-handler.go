package handler

import (
	"crud_api_company/controller/employee"
	"crud_api_company/utils"
	"net/http"
	"strconv"

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

func (h *handlerEmployee) AddEmployeeIntoCompany(ctx *gin.Context) {

	companyParam := ctx.Param("companyId")

	companyParamConv, err := strconv.Atoi(companyParam)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Add failed", http.StatusBadRequest, http.MethodPost, err)
		return
	}

	var input employee.EmployeeCreateInput
	ctx.ShouldBindJSON(&input)

	result, erro := h.service.AddEmployeeIntoCompany(companyParamConv, &input)

	if erro != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Add failed", http.StatusBadRequest, http.MethodPost, erro)
		return
	}

	utils.APIResponse(ctx, "Add succ", http.StatusOK, http.MethodPost, result)
}

func (h *handlerEmployee) DeleteEmployee(ctx *gin.Context) {

	employeeParam := ctx.Param("employeeId")

	employeeParamConv, err := strconv.Atoi(employeeParam)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Delete failed", http.StatusBadRequest, http.MethodDelete, err)
		return
	}

	result, errD := h.service.DeleteEmployee(employeeParamConv)

	if errD != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Delete failed", http.StatusBadRequest, http.MethodDelete, errD)
		return
	}

	utils.APIResponse(ctx, "Delete Successfully", http.StatusOK, http.MethodDelete, result)
}

func (h *handlerEmployee) GetAllEmployeeFromCompany(ctx *gin.Context) {
	companyParam := ctx.Param("companyId")

	companyParamConv, err := strconv.Atoi(companyParam)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Get failed", http.StatusBadRequest, http.MethodGet, err)
		return
	}

	result, errG := h.service.GetAllEmployeeFromCompany(companyParamConv)

	if errG != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get failed", http.StatusBadRequest, http.MethodGet, errG)
		return
	}

	utils.APIResponse(ctx, "Get Successfully", http.StatusOK, http.MethodGet, result)
}
