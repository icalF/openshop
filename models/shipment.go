package models

import "time"

const (
	DELIVERED             = "DELIVERED"
	WAITING_TO_BE_SHIPPED = "WAITING_TO_BE_SHIPPED"
	SHIPPED               = "SHIPPED"
)

type Shipment struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Name      string `json:"name" validate:"required,max=40,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	OrderID   int64  `json:"order_id"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
