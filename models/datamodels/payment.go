package datamodels

import "time"

const (
	WAITING_PAYMENT = "WAITING_PAYMENT"
	TO_BE_VERIFIED  = "TO_BE_VERIFIED"
	VERIFIED        = "VERIFIED"
)

type Payment struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	OrderID   int64     `json:"order_id"`
	Amount    int       `json:"amount"`
	Proof     string    `json:"proof"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPayment(orderID int64, netAmount int) Payment {
	return Payment{
		OrderID: orderID,
		Amount:  netAmount,
		Status:  WAITING_PAYMENT,
	}
}
