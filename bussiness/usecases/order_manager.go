package usecases

import "github.com/koneko096/openshop/models/datamodels"

type OrderManager interface {
	GetAll() []datamodels.Order
	GetAllSubmitted() []datamodels.Order
	GetByID(id int64) (datamodels.Order, bool)
	GetByUserID(userId int64) []datamodels.Order
	InsertOrUpdate(order datamodels.Order) (datamodels.Order, error)
	DeleteByID(id int64) bool
}
