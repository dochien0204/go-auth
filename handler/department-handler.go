package handler

import (
	"crud_api_company/controller/department"
	"crud_api_company/utils"
	"fmt"
	"net/http"
	"strconv"

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

func (h *handlerDepartment) GetAllDepartmentHandler(ctx *gin.Context) {

	result, err := h.service.GetAllDepartMent()

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get All Department Failed", http.StatusBadRequest, http.MethodGet, err)
		return
	}
	fmt.Println("Nhu cc")

	utils.APIResponse(ctx, "Get all department successfully", http.StatusOK, http.MethodGet, result)
}

func (h *handlerDepartment) CreateNewDepartmetForCompany(ctx *gin.Context) {
	var input department.DepartmentCreateInput
	ctx.ShouldBindJSON(&input)

	result, err := h.service.CreateNewDepartmetForCompany(&input)

	if err != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Create Failed", http.StatusBadRequest, http.MethodPost, err)
		return
	}

	utils.APIResponse(ctx, "Create New Department Successfully", http.StatusOK, http.MethodPost, result)

}

func (h *handlerDepartment) GetDepartmentById(ctx *gin.Context) {

	idPath := ctx.Param("departmentId")

	convertId, err := strconv.Atoi(idPath)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusBadRequest, http.MethodGet, "Path is invalid")
		return
	}

	result, failed := h.service.GetDepartmentById(convertId)

	if failed != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusBadRequest, http.MethodGet, failed)
		return
	}

	utils.APIResponse(ctx, "Get Department Successfully", http.StatusOK, http.MethodGet, result)
}

func (h *handlerDepartment) UpdateDepartmentById(ctx *gin.Context) {

	pathVariable := ctx.Param("departmentId")

	pathConv, err := strconv.Atoi(pathVariable)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusBadRequest, http.MethodPatch, "Path is invalid")
		return
	}

	var input department.DepartmentUpdateInput
	ctx.ShouldBindJSON(&input)

	result, failed := h.service.UpdateDepartmentById(pathConv, &input)

	if failed != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Update Failed", http.StatusBadRequest, http.MethodPatch, failed)
		return
	}

	utils.APIResponse(ctx, "Update Department Successfully", http.StatusOK, http.MethodPatch, result)
}

func (h *handlerDepartment) DeleteDepartmentById(ctx *gin.Context) {
	pathVariable := ctx.Param("departmentId")

	pathConv, err := strconv.Atoi(pathVariable)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusBadRequest, http.MethodDelete, "Path is invalid")
		return
	}

	result, errDel := h.service.DeleteDepartmentById(pathConv)
	if errDel != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Update Failed", http.StatusBadRequest, http.MethodDelete, errDel)
		return
	}

	utils.APIResponse(ctx, "Update Department Successfully", http.StatusOK, http.MethodDelete, result)

}

func (h *handlerDepartment) GetAllEmployeeInDepartment(ctx *gin.Context) {
	pathVariable := ctx.Param("departmentId")

	pathConv, err := strconv.Atoi(pathVariable)

	if err != nil {
		utils.ValidatorErrorResponse(ctx, "Get Failed", http.StatusBadRequest, http.MethodDelete, "Path is invalid")
		return
	}

	result, errGetAll := h.service.GetAllEmployeeInDepartment(pathConv)

	if errGetAll != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Update Failed", http.StatusBadRequest, http.MethodGet, errGetAll)
		return
	}

	utils.APIResponse(ctx, "Get All Employees in Department Successfully", http.StatusOK, http.MethodGet, result)
}

func (h *handlerDepartment) DeleteEmployeeInDepartment(ctx *gin.Context) {
	employeeIdParam := ctx.Param("employeeId")
	departmentIdParam := ctx.Param("departmentId")

	employeeIdConv, errE := strconv.Atoi(employeeIdParam)
	departmentIdConv, errD := strconv.Atoi(departmentIdParam)

	if errE != nil || errD != nil {
		utils.ValidatorErrorResponse(ctx, "Delete Employee Failed", http.StatusBadRequest, http.MethodDelete, "Path invalid")
		return
	}

	result, errorD := h.service.DeleteEmployeeInDepartment(employeeIdConv, departmentIdConv)

	if errorD != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Delete Employee Failed", http.StatusBadRequest, http.MethodDelete, errorD)
		return
	}

	utils.APIResponse(ctx, "Delete Successfully", http.StatusOK, http.MethodDelete, result)
}

func (h *handlerDepartment) AddEmployeeIntoDepartment(ctx *gin.Context) {
	employeeIdParam := ctx.Param("employeeId")
	departmentIdParam := ctx.Param("departmentId")

	employeeIdConv, errE := strconv.Atoi(employeeIdParam)
	departmentIdConv, errD := strconv.Atoi(departmentIdParam)

	if errE != nil || errD != nil {
		utils.ValidatorErrorResponse(ctx, "Delete Employee Failed", http.StatusBadRequest, http.MethodDelete, "Path invalid")
		return
	}

	result, errGet := h.service.AddEmployeeIntoDepartment(employeeIdConv, departmentIdConv)

	if errGet != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "POST Failed", http.StatusBadRequest, http.MethodPost, errGet)
		return
	}

	utils.APIResponse(ctx, "ADD Successfully", http.StatusOK, http.MethodPost, result)

}
