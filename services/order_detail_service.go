package services

import (
	"github.com/icalF/openshop/dao"
	"github.com/icalF/openshop/models/datamodels"
)

type OrderDetailService interface {
	GetAll() []datamodels.OrderDetail
	GetByID(id int64) (datamodels.OrderDetail, bool)
	InsertOrUpdate(orderDetail datamodels.OrderDetail) (datamodels.OrderDetail, error)
	DeleteByID(id int64) bool
}

func NewOrderDetailService(dao dao.OrderDetailDAO) OrderDetailService {
	return &orderDetailService{
		dao: dao,
	}
}

type orderDetailService struct {
	dao dao.OrderDetailDAO
}

func (s *orderDetailService) GetAll() []datamodels.OrderDetail {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *orderDetailService) GetByID(id int64) (datamodels.OrderDetail, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *orderDetailService) InsertOrUpdate(orderDetail datamodels.OrderDetail) (datamodels.OrderDetail, error) {
	return s.dao.InsertOrUpdate(orderDetail)
}

func (s *orderDetailService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}