package repo

import (
	"BackEnd/models"
	"gorm.io/gorm"
)

// 获取用户根据地址
func GetUserByAddress(db *gorm.DB, address string) (*models.User, error) {
	var user models.User
	if err := db.Where("address = ?", address).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 没找到用户
		}
		return nil, err // 其他错误
	}
	return &user, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}
