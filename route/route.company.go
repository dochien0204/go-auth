package route

import (
	"crud_api_company/controller/company"
	"crud_api_company/controller/department"
	"crud_api_company/handler"

	"github.com/gin-gonic/gin"
)

func InItCompanyRoute(route *gin.Engine) {

	companyGroup := route.Group("/api/v1")
	companyRepository := company.NewRepository()
	companyService := company.NewService(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyService)

	departmentRepository := department.NewDepartmentRepository()
	departmentService := department.NewDepartmentService(departmentRepository)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	companyGroup.GET("/company/list", companyHandler.GetAllCompanyHanlder)

	//department
	companyGroup.GET("/department/list", departmentHandler.GetAllDepartmentHandler)
	companyGroup.POST("/department/create", departmentHandler.CreateNewDepartmetForCompany)
	companyGroup.GET("/department/:departmentId", departmentHandler.GetDepartmentById)
	companyGroup.PATCH("/department/:departmentId/update", departmentHandler.UpdateDepartmentById)
	companyGroup.DELETE("/department/:departmentId/delete", departmentHandler.DeleteDepartmentById)
	companyGroup.GET("/department/:departmentId/get-employees", departmentHandler.GetAllEmployeeInDepartment)
}
