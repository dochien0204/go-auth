package employee

type EmployeeCreateInput struct {
	EmployeeName  string `gorm:"size:255;not null" json:"employeeName"`
	EmployeeAge   int    `json:"employeeAge"`
	EmployeePhone string `gorm:"size:12" json:"employeePhone"`
}
