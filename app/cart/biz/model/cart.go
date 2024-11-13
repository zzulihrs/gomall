package model

import (
	"context"
	"errors"
	"fmt"
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
	//// 添加到购物车
	//if cart.UserId == 0 || cart.ProductId == 0 || cart.Qty == 0 {
	//	return errors.New("userId productId qty  is required")
	//}
	//
	//err := db.Create(&cart).Error
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//// 使用 Create 方法创建记录
	//return err
	var row Cart

	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
		First(&row).Error

	if err == nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID != 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", cart.Qty)).Error
	}
	return db.WithContext(ctx).Model(&Cart{}).Create(cart).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Delete(&Cart{}).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) (cartList []*Cart, err error) {
	if db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	err = db.Debug().
		WithContext(ctx).
		Model(&Cart{}).
		Find(&cartList, "user_id = ?", userId).Error
	fmt.Println("cartList", cartList)
	return cartList, err
}
