package router

import (
	"BackEnd/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 跨域中间件（前后端分离时需要）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 注册接口
	r.POST("/api/register", api.RegisterHandler(db))
	// 登录接口
	r.POST("/api/login", api.LoginWithMetaMask)
	r.POST("/api/product", api.CreateProduct)
	r.POST("/api/logistics", api.CreateLogistics)
	r.GET("/api/getAllProducts", api.GetAllProducts) // 获取所有产品
	r.GET("/api/getAllLogistics", api.GetAllLogistics)
	r.POST("/api/updateProductStatus", api.UpdateProductStatus)
	r.POST("/api/updateShippingStatus", api.UpdateLogisticsStatus)
	r.POST("/api/GetProductByBatchId", api.GetProductByBatchId)
	r.POST("/api/GetLogisticsByBatchId", api.GetLogisticsByBatchId)
	return r
}
