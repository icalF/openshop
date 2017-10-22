package models

import "time"

type Product struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Name      string `json:"name" validate:"required,max=40,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Price     int    `json:"price" validate:"required,min=500"`
	Qty       int    `json:"qty" validate:"required,min=0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
