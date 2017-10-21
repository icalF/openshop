package models

import "time"

type User struct {
	ID     int64  `json:"id", gorm:"primary_key"`
	Name   string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
