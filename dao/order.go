package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type OrderDAO interface {
	Select(query Query) (model models.Order, found bool)
	SelectMany(query Query, limit int) (results []models.Order)

	Insert(model models.Order) (models.Order, error)
	Delete(query Query) (deleted bool)
}

type orderRepository struct {
	source *gorm.DB
}

func NewOrderDAO(connection *gorm.DB) OrderDAO {
	return &orderRepository{source: connection}
}

func (r *orderRepository) Select(query Query) (models.Order, bool) {
	order := models.Order{}
	if err := r.source.Where(query).First(&order).Error; err != nil {
		return models.Order{}, false
	}
	return order, true
}

func (r *orderRepository) SelectMany(query Query, limit int) (results []models.Order) {
	orders := new([]models.Order)
	r.source.Where(query).Find(&orders).Limit(limit)
	return *orders
}

func (r *orderRepository) Insert(order models.Order) (models.Order, error) {
	err := r.source.Create(&order).Error
	return order, err
}

func (r *orderRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.Order{}, query).Error; err != nil {
		return false
	}
	return true
}
