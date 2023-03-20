package utils

import "gorm.io/gorm"

func CheckError(result *gorm.DB) (bool, string) {
	if result.Error != nil {
		return false, "Something went wrong"
	}

	if result.RowsAffected == 0 {
		return false, "Not exists"
	}

	return true, "Successfully"
}
