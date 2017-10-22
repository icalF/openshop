package datamodels

import "time"

const (
	UNSUBMITTED = "UNSUBMITTED"
	SUBMITTED   = "SUBMITTED"
	FINISHED    = "FINISHED"
)

type Order struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	UserID    int64     `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrder(userID int64) Order {
	return Order{
		UserID: userID,
		Status: UNSUBMITTED,
	}
}
