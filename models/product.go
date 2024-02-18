package models

type Product struct {
	id int64 `gorm:"primaryKey" json:"id"`
}
