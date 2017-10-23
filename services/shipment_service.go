package services

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type ShipmentService interface {
	GetAll() []datamodels.Shipment
	GetByID(id int64) (datamodels.Shipment, bool)
	GetByOrderID(orderId int64) (datamodels.Shipment, bool)
	GetByShippingCode(code string) (datamodels.Shipment, bool)
	InsertOrUpdate(shipment datamodels.Shipment) (datamodels.Shipment, error)
	DeleteByID(id int64) bool
}

func NewShipmentService(dao dao.ShipmentDAO) ShipmentService {
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