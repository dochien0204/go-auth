package handler

import (
	"crud_api_company/controller/company"
	"crud_api_company/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerCompany struct {
	service company.ICompany
}

func NewCompanyHandler(service company.ICompany) *handlerCompany {
	return &handlerCompany{
		service: service,
	}
}

func (h *handlerCompany) GetAllCompanyHanlder(ctx *gin.Context) {

	listCompany, errGetAll := h.service.GetAllCompany()

	if errGetAll != "Successfully" {
		utils.ValidatorErrorResponse(ctx, "Get All Companies Failed", http.StatusBadRequest, http.MethodGet, errGetAll)
		return
	}

	utils.APIResponse(ctx, "Get All Companies Successfully", http.StatusOK, http.MethodGet, listCompany)

}
