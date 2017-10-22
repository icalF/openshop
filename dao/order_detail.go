package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type OrderDetailDAO interface {
	Select(query Query) (model models.OrderDetail, found bool)
	SelectMany(query Query, limit int) (results []models.OrderDetail)

	InsertOrUpdate(model models.OrderDetail) (models.OrderDetail, error)
	Delete(query Query) (deleted bool)
}

type orderDetailRepository struct {
	source *gorm.DB
}

func NewOrderDetailDAO(connection *gorm.DB) OrderDetailDAO {
	return &orderDetailRepository{source: connection}
}

func (r *orderDetailRepository) Select(query Query) (models.OrderDetail, bool) {
	orderDetail := models.OrderDetail{}
	if err := r.source.Where(query).First(&orderDetail).Error; err != nil {
		return models.OrderDetail{}, false
	}
	return orderDetail, true
}

func (r *orderDetailRepository) SelectMany(query Query, limit int) (results []models.OrderDetail) {
	orderDetails := new([]models.OrderDetail)
	r.source.Where(query).Find(&orderDetails).Limit(limit)
	return *orderDetails
}

func (r *orderDetailRepository) InsertOrUpdate(orderDetail models.OrderDetail) (_ models.OrderDetail, err error) {
	var oldOrderDetail models.OrderDetail
	if err := r.source.First(&oldOrderDetail).Error; err != nil {
		r.source.Create(&orderDetail)
	} else {
		r.source.Model(&oldOrderDetail).Update(&orderDetail)
	}

	return orderDetail, err
}

func (r *orderDetailRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.OrderDetail{}, query).Error; err != nil {
		return false
	}
	return true
}
