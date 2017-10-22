package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type ShipmentService interface {
	GetAll() []models.Shipment
	GetByID(id int64) (models.Shipment, bool)
	InsertOrUpdate(shipment models.Shipment) (models.Shipment, error)
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

func (s *shipmentService) GetAll() []models.Shipment {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *shipmentService) GetByID(id int64) (models.Shipment, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *shipmentService) InsertOrUpdate(shipment models.Shipment) (models.Shipment, error) {
	return s.dao.InsertOrUpdate(shipment)
}

func (s *shipmentService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}