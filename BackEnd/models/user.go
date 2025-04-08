package models

type User struct {
	ID      uint   `gorm:"primaryKey"`
	Address string `gorm:"uniqueIndex"`
	Role    string
}
