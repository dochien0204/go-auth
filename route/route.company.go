package route

import (
	"crud_api_company/controller/company"
	"crud_api_company/controller/department"
	"crud_api_company/controller/employee"
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

	employeeRepository := employee.NewEmployeeRepository()
	employeeService := employee.NewEmployeeService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	companyGroup.GET("/company/list", companyHandler.GetAllCompanyHanlder)

	//department
	companyGroup.GET("/department/list", departmentHandler.GetAllDepartmentHandler)
	companyGroup.POST("/department/create", departmentHandler.CreateNewDepartmetForCompany)
	companyGroup.GET("/department/:departmentId", departmentHandler.GetDepartmentById)
	companyGroup.PATCH("/department/:departmentId/update", departmentHandler.UpdateDepartmentById)
	companyGroup.DELETE("/department/:departmentId/delete", departmentHandler.DeleteDepartmentById)
	companyGroup.GET("/department/:departmentId/get-employees", departmentHandler.GetAllEmployeeInDepartment)
	companyGroup.DELETE("/department/delete-employee-department/:employeeId/:departmentId", departmentHandler.DeleteEmployeeInDepartment)
	companyGroup.POST("/department/add-employee-into-department/:employeeId/:departmentId", departmentHandler.AddEmployeeIntoDepartment)

	//employee
	companyGroup.GET("/employee/list", employeeHandler.GetAllEmployee)
	companyGroup.POST("/employee/add-employee-into-company/:companyId", employeeHandler.AddEmployeeIntoCompany)
	companyGroup.DELETE("/employee/:employeeId/delete", employeeHandler.DeleteEmployee)
	companyGroup.GET("/employee/:companyId/get", employeeHandler.GetAllEmployeeFromCompany)
}
