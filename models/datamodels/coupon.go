package datamodels

import "time"

type Coupon struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Code      string `json:"code" validate:"required,len=6,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Percent   int    `json:"percent" validate:"required,max=100"`
	Nominal   int    `json:"nominal" validate:"required,min=500"`
	Qty       int    `json:"qty" validate:"required,min=0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
