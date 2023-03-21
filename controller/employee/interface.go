package employee

import "crud_api_company/model"

type IEmployee interface {
	GetAllEmployee() ([]model.Employee, string)
}
