package service

import (
	"BackEnd/db"
	"BackEnd/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// 更新物流的审核状态和交易哈希
// UpdateLogisticsStatus 更新物流审核状态和交易哈希
func UpdateLogisticsStatus(id int, status string, txHash string) error {
	// 获取数据库连接
	dbConn := db.GetDB()

	// 更新物流信息
	result := dbConn.Model(&models.Logistics{}).Where("id = ?", id).Updates(models.Logistics{
		Status: status,
		TxHash: txHash,
	})

	if result.Error != nil {
		return fmt.Errorf("更新物流状态失败: %v", result.Error)
	}

	return nil
}

// 获取所有物流数据接口
// 获取所有物流数据接口
// 获取所有物流数据接口
func GetAllLogistics() ([]models.Logistics, error) {
	var logistics []models.Logistics
	err := db.GetDB().Find(&logistics).Error
	if err != nil {
		return nil, err
	}
	return logistics, nil
}

// 获取物流数据接口
func GetLogisticsByBatchId(batchId string) ([]models.Logistics, error) {
	var logistics []models.Logistics
	if err := db.GetDB().Where("batch_id = ?", batchId).Find(&logistics).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("物流数据未找到")
		}
		return nil, err
	}
	return logistics, nil
}
