package usecases

import "github.com/koneko096/openshop/models/datamodels"

type OrderDetailManager interface {
	GetAll() []datamodels.OrderDetail
	GetByID(id int64) (datamodels.OrderDetail, bool)
	GetByOrderID(id int64) []datamodels.OrderDetail
	InsertOrUpdate(orderDetail datamodels.OrderDetail) (datamodels.OrderDetail, error)
	DeleteByID(id int64) bool
}
