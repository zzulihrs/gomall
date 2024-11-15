package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32  `gorm:"type:int(11)"`
	OrderIdRefer string  `gorm:"size:256;index"`
	Quantity     uint32  `gorm:"type:int(11)"`
	Cost         float32 `gorm:"type:decimal(10, 2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
