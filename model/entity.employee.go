package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeName  string        `gorm:"size:255;not null" json:"employeeName"`
	EmployeeAge   int           `json:"employeeAge"`
	EmployeePhone string        `gorm:"size:12" json:"employeePhone"`
	Departments   []*Department `gorm:"many2many:employee_department;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
