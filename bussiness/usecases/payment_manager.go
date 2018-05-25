package usecases

import "github.com/koneko096/openshop/models/datamodels"

type PaymentManager interface {
	GetAll() []datamodels.Payment
	GetByID(id int64) (datamodels.Payment, bool)
	GetByOrderID(orderId int64) (datamodels.Payment, bool)
	InsertOrUpdate(payment datamodels.Payment) (datamodels.Payment, error)
	DeleteByID(id int64) bool
}
