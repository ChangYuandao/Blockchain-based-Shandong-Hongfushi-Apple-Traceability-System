package api

import (
	"BackEnd/db"
	"BackEnd/models"
	"BackEnd/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询物流信息接口
// 查询物流信息接口
func GetLogisticsByBatchId(c *gin.Context) {
	var request struct {
		BatchId string `json:"batchId"` // 接收前端上传的 batchId
	}

	// 解析请求体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务层获取物流数据
	logistics, err := service.GetLogisticsByBatchId(request.BatchId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 返回查询到的数据
	c.JSON(http.StatusOK, gin.H{"data": logistics})
}

func UpdateLogisticsStatus(c *gin.Context) {
	var data struct {
		Id     int    `json:"id"`
		Status string `json:"status"`
		TxHash string `json:"txHash"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := service.UpdateLogisticsStatus(data.Id, data.Status, data.TxHash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "物流状态更新成功"})
}

func CreateLogistics(c *gin.Context) {
	var logistics models.Logistics
	if err := c.ShouldBindJSON(&logistics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.GetDB().Create(&logistics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存物流信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "物流信息创建成功"})
}

// 获取所有物流数据接口
func GetAllLogistics(c *gin.Context) {
	// 调用 service 层的 GetAllLogistics 函数
	logistics, err := service.GetAllLogistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取物流数据失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": logistics})
}
