package models

import "time"

const (
	SUBMITTED = "SUBMITTED"
	FINISHED  = "FINISHED"
)

type Order struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	UserID    int64 `json:"order_id"`
	Status    int64 `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
