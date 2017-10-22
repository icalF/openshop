package models

import "time"

const (
	WAITING_PAYMENT = "WAITING_PAYMENT"
	TO_BE_VERIFIED = "TO_BE_VERIFIED"
	VERIFIED = "VERIFIED"
)

type Payment struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	OrderID   int64  `json:"order_id"`
	Amount    int    `json:"amount"`
	Bank	  string `json:"bank"`
	Proof     string `json:"proof"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
