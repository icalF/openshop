package usecases

import "github.com/koneko096/openshop/models/datamodels"

type ShipmentManager interface {
	GetAll() []datamodels.Shipment
	GetByID(id int64) (datamodels.Shipment, bool)
	GetByOrderID(orderId int64) (datamodels.Shipment, bool)
	GetByShippingCode(code string) (datamodels.Shipment, bool)
	InsertOrUpdate(shipment datamodels.Shipment) (datamodels.Shipment, error)
	DeleteByID(id int64) bool
}
