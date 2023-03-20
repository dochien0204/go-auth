package company

import "crud_api_company/model"

type ICompany interface {
	GetAllCompany() ([]model.Company, string)
}
