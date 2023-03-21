package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string      `gorm:"size:255" json:"departmentName"`
	CompanyId      uint        `gorm:"not null"`
	Employees      []*Employee `gorm:"many2many:employee_department"`
}
