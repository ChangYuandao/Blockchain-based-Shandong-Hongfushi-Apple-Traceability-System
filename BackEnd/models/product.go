package models

type Product struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ProductName string `json:"productName"`
	BatchId     string `json:"batchId"`
	BaseName    string `json:"baseName"`
	Location    string `json:"location"`
	PlantDate   string `json:"plantDate"`
	HarvestDate string `json:"harvestDate"`
	SoilType    string `json:"soilType"`
	Status      string `json:"status"`        // 新增字段，表示审核状态
	TxHash      string `gorm:"column:txHash"` // 注意大小写完全匹配数据库
}
