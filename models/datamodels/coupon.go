package datamodels

import "time"

type Coupon struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Code      string    `json:"code" gorm:"unique" validate:"len=6,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Percent   int       `json:"percent" validate:"max=100"`
	Nominal   int       `json:"nominal" validate:"min=0"`
	Due       time.Time `json:"due"`
	Qty       int       `json:"qty" validate:"required,min=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
