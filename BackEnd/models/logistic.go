package models

type Logistics struct {
	ID                   uint   `gorm:"primaryKey" json:"id"`
	BatchId              string `json:"batchId"`
	ShippingCompany      string `json:"shippingCompany"`
	TrackingNumber       string `json:"trackingNumber"`
	ShippingStatus       string `json:"shippingStatus"`
	EstimatedArrivalDate string `json:"estimatedArrivalDate"` // 格式：2025-04-10
	ActualArrivalDate    string `json:"actualArrivalDate"`    // 格式：2025-04-12
	ShippingMethod       string `json:"shippingMethod"`
	ShippingCost         string `json:"shippingCost"`  // 示例："120.50"
	Status               string `json:"status"`        // 新增字段，表示审核状态
	TxHash               string `gorm:"column:txHash"` // 注意大小写完全匹配数据库
}
