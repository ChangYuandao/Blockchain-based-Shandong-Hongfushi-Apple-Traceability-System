package service

import (
	"BackEnd/db"
	"BackEnd/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// 更新产品状态
func UpdateProductStatus(id int, status string, txHash string) error {
	// 获取数据库连接
	dbConn := db.GetDB()

	// 更新产品信息
	result := dbConn.Model(&models.Product{}).Where("id = ?", id).Updates(models.Product{
		Status: status,
		TxHash: txHash,
	})

	if result.Error != nil {
		return fmt.Errorf("更新产品状态失败: %v", result.Error)
	}

	return nil
}
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.GetDB().Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// 获取单条产品数据接口
// 获取产品数据接口
func GetProductByBatchId(batchId string) ([]models.Product, error) {
	var product []models.Product
	if err := db.GetDB().Where("batch_id = ?", batchId).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("产品数据未找到")
		}
		return nil, err
	}
	return product, nil
}
