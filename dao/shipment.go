package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type ShipmentDAO interface {
	Select(query Query) (model models.Shipment, found bool)
	SelectMany(query Query, limit int) (results []models.Shipment)

	InsertOrUpdate(model models.Shipment) (models.Shipment, error)
	Delete(query Query) (deleted bool)
}

type shipmentRepository struct {
	source *gorm.DB
}

func NewShipmentDAO(connection *gorm.DB) ShipmentDAO {
	return &shipmentRepository{source: connection}
}

func (r *shipmentRepository) Select(query Query) (models.Shipment, bool) {
	shipment := models.Shipment{}
	if err := r.source.Where(query).First(&shipment).Error; err != nil {
		return models.Shipment{}, false
	}
	return shipment, true
}

func (r *shipmentRepository) SelectMany(query Query, limit int) (results []models.Shipment) {
	shipments := new([]models.Shipment)
	r.source.Where(query).Find(&shipments).Limit(limit)
	return *shipments
}

func (r *shipmentRepository) InsertOrUpdate(shipment models.Shipment) (_ models.Shipment, err error) {
	var oldShipment models.Shipment
	if err := r.source.First(&oldShipment).Error; err != nil {
		r.source.Create(&shipment)
	} else {
		r.source.Model(&oldShipment).Update(&shipment)
	}

	return shipment, err
}

func (r *shipmentRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.Shipment{}, query).Error; err != nil {
		return false
	}
	return true
}
