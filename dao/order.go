package dao

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type OrderDAO interface {
	Select(query Query) (model datamodels.Order, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Order)

	Insert(model datamodels.Order) (datamodels.Order, error)
	Delete(query Query) (deleted bool)
}

type orderRepository struct {
	source *gorm.DB
}

func NewOrderDAO(connection *gorm.DB) OrderDAO {
	return &orderRepository{source: connection}
}

func (r *orderRepository) Select(query Query) (datamodels.Order, bool) {
	order := datamodels.Order{}
	if err := r.source.Where(query).First(&order).Error; err != nil {
		return datamodels.Order{}, false
	}
	return order, true
}

func (r *orderRepository) SelectMany(query Query, limit int) (results []datamodels.Order) {
	orders := new([]datamodels.Order)
	r.source.Where(query).Find(&orders).Limit(limit)
	return *orders
}

func (r *orderRepository) Insert(order datamodels.Order) (datamodels.Order, error) {
	err := r.source.Create(&order).Error
	return order, err
}

func (r *orderRepository) Delete(query Query) bool {
	err := r.source.Delete(datamodels.Order{}, query).Error
	return err == nil
}
