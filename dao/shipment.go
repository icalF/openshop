package dao

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type ShipmentDAO interface {
	Select(query Query) (model datamodels.Shipment, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Shipment)

	InsertOrUpdate(model datamodels.Shipment) (datamodels.Shipment, error)
	Delete(query Query) (deleted bool)
}

type shipmentRepository struct {
	source *gorm.DB
}

func NewShipmentDAO(connection *gorm.DB) ShipmentDAO {
	return &shipmentRepository{source: connection}
}

func (r *shipmentRepository) Select(query Query) (datamodels.Shipment, bool) {
	shipment := datamodels.Shipment{}
	if err := r.source.Where(query).First(&shipment).Error; err != nil {
		return datamodels.Shipment{}, false
	}
	return shipment, true
}

func (r *shipmentRepository) SelectMany(query Query, limit int) (results []datamodels.Shipment) {
	shipments := new([]datamodels.Shipment)
	r.source.Where(query).Find(&shipments).Limit(limit)
	return *shipments
}

func (r *shipmentRepository) InsertOrUpdate(shipment datamodels.Shipment) (_ datamodels.Shipment, err error) {
	var oldShipment datamodels.Shipment
	if err := r.source.First(&oldShipment).Error; err != nil {
		r.source.Create(&shipment)
	} else {
		r.source.Model(&oldShipment).Update(&shipment)
	}

	return shipment, err
}

func (r *shipmentRepository) Delete(query Query) bool {
	if err := r.source.Delete(datamodels.Shipment{}, query).Error; err != nil {
		return false
	}
	return true
}
