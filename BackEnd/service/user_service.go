package service

import (
	"BackEnd/db"
	"BackEnd/models"
	"BackEnd/repo"
	"errors"
	"gorm.io/gorm"
)

// 根据 MetaMask 地址获取用户角色
func GetUserRole(address string) (string, error) {
	user, err := repo.GetUserByAddress(db.GetDB(), address) // 传递 db.DB 实例
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", nil // 用户不存在
	}
	return user.Role, nil
}

func RegisterUser(db *gorm.DB, address, role string) error {
	// 检查是否已存在
	user, err := repo.GetUserByAddress(db, address)
	if err != nil {
		return err // 如果查询出错，直接返回错误
	}
	if user != nil {
		return errors.New("该地址已注册")
	}

	// 插入
	newUser := &models.User{
		Address: address,
		Role:    role,
	}
	return repo.CreateUser(db, newUser)
}
