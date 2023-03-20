package department

import "crud_api_company/model"

type IDepartment interface {
	GetAllDepartMent() ([]model.Department, string)
}
