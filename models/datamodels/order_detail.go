package datamodels

import "time"

type OrderDetail struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	OrderID   int64     `json:"order_id"`
	ProductID int64     `json:"product_id"`
	Qty       int64     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrderDetail(orderID int64) OrderDetail {
	return OrderDetail{
		OrderID: orderID,
		Qty:     1,
	}
}
