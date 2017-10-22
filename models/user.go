package models

import "time"

type User struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Name      string `json:"name" validate:"required,max=40,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Email     string `json:"email" validate:"required,max=40,email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
