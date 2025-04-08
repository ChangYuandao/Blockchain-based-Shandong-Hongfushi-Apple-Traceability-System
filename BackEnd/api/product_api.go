package api

import (
	"BackEnd/db"
	"BackEnd/models"
	"BackEnd/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductByBatchId(c *gin.Context) {
	var request struct {
		BatchId string `json:"batchId"` // 接收前端上传的 batchId
	}

	// 解析请求体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用服务层获取产品数据
	product, err := service.GetProductByBatchId(request.BatchId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 返回查询到的数据
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// 更新产品审核状态的 API
func UpdateProductStatus(c *gin.Context) {
	var data struct {
		Id     int    `json:"id"`
		Status string `json:"status"`
		TxHash string `json:"txHash"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := service.UpdateProductStatus(data.Id, data.Status, data.TxHash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "产品状态更新成功"})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将数据保存到数据库
	if err := db.GetDB().Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// 获取所有产品数据接口
func GetAllProducts(c *gin.Context) {
	// 调用 service 层的 GetAllProducts 函数
	products, err := service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取产品数据失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}
