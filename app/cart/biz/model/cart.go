package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       uint32 `json:"qty"`
}

func (c Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, cart *Cart) error {
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
		First(&row).Error
	if err == nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
			Update("qty", row.Qty+cart.Qty).Error
	}
	return db.WithContext(ctx).Model(&Cart{}).Create(cart).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Delete(&Cart{}).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&rows).Error
	return rows, err
}
