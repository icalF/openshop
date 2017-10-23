package datamodels

import "time"

const (
	DELIVERED = "DELIVERED"
	SHIPPED   = "SHIPPED"
)

type Shipment struct {
	ID           int64  `json:"id" gorm:"primary_key"`
	ShippingCode string `json:"name" validate:"required,len=8,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	OrderID      int64  `json:"order_id"`
	Status       string `json:"status"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewShipment(orderID int64, code string) Shipment {
	return Shipment{
		OrderID:      orderID,
		ShippingCode: code,
		Status:       SHIPPED,
	}
}