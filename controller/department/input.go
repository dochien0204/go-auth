package department

type DepartmentCreateInput struct {
	CompanyId      int    `json:"companyId"`
	DepartmentName string `json:"departmentName"`
}

type DepartmentUpdateInput struct {
	CompanyId      int    `json:"companyId"`
	DepartmentName string `json:"departmentName"`
}
