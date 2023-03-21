package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName   string `gorm:"size:255;not null;" json:"companyName"`
	TotalEmployee int    `json:"totalEmployee"`
	Departments   []Department
	Employees     []Employee
}
