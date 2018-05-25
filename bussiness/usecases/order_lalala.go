package usecases

import "github.com/koneko096/openshop/models/datamodels"

type OrderLalala interface {
	InsertCoupon(id int64, code string) (bool, error)
	Checkout(id int64) (datamodels.Payment, bool, error)
	GetTotalAmount(order datamodels.Order) int
}
