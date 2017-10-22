package datamodels

import "time"

type Product struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Name      string `json:"name" validate:"max=40,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Price     int    `json:"price" validate:"min=0"`
	Qty       int    `json:"qty" validate:"min=0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
