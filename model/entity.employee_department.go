package model

type EmployeeDepartment struct {
	EmployeeId   int `gorm:"primaryKey"`
	DepartmentId int `gorm:"primaryKey;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
}
