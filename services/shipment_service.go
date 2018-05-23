package services

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/dao"
)

func NewShipmentManager(dao dao.ShipmentDAO) ShipmentManager {
	return &shipmentService{
		dao: dao,
	}
}


type shipmentService struct {
	dao dao.ShipmentDAO
}



func (s *shipmentService) GetAll() []datamodels.Shipment {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *shipmentService) GetByID(id int64) (datamodels.Shipment, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *shipmentService) GetByOrderID(orderId int64) (datamodels.Shipment, bool) {
	return s.dao.Select(map[string]string{
		"order_id": string(orderId),
	})
}

func (s *shipmentService) GetByShippingCode(code string) (datamodels.Shipment, bool) {
	return s.dao.Select(map[string]string{
		"shipping_code": code,
	})
}

func (s *shipmentService) InsertOrUpdate(shipment datamodels.Shipment) (datamodels.Shipment, error) {
	return s.dao.InsertOrUpdate(shipment)
}

func (s *shipmentService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}