package dao

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type OrderDetailDAO interface {
	Select(query Query) (model datamodels.OrderDetail, found bool)
	SelectMany(query Query, limit int) (results []datamodels.OrderDetail)

	InsertOrUpdate(model datamodels.OrderDetail) (datamodels.OrderDetail, error)
	Delete(query Query) (deleted bool)
}

type orderDetailRepository struct {
	source *gorm.DB
}

func NewOrderDetailDAO(connection *gorm.DB) OrderDetailDAO {
	return &orderDetailRepository{source: connection}
}

func (r *orderDetailRepository) Select(query Query) (datamodels.OrderDetail, bool) {
	orderDetail := datamodels.OrderDetail{}
	if err := r.source.Where(query).First(&orderDetail).Error; err != nil {
		return datamodels.OrderDetail{}, false
	}
	return orderDetail, true
}

func (r *orderDetailRepository) SelectMany(query Query, limit int) (results []datamodels.OrderDetail) {
	orderDetails := new([]datamodels.OrderDetail)
	r.source.Where(query).Find(&orderDetails).Limit(limit)
	return *orderDetails
}

func (r *orderDetailRepository) InsertOrUpdate(orderDetail datamodels.OrderDetail) (_ datamodels.OrderDetail, err error) {
	var oldOrderDetail datamodels.OrderDetail
	if err := r.source.First(&oldOrderDetail, orderDetail.ID).Error; err != nil {
		r.source.Create(&orderDetail)
	} else {
		r.source.Model(&oldOrderDetail).Update(&orderDetail)
	}

	return orderDetail, err
}

func (r *orderDetailRepository) Delete(query Query) bool {
	err := r.source.Delete(datamodels.Order{}, query).Error
	return err == nil
}
