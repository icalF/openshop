package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type OrderDetailService interface {
	GetAll() []models.OrderDetail
	GetByID(id int64) (models.OrderDetail, bool)
	InsertOrUpdate(orderDetail models.OrderDetail) (models.OrderDetail, error)
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

func (s *orderDetailService) GetAll() []models.OrderDetail {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *orderDetailService) GetByID(id int64) (models.OrderDetail, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *orderDetailService) InsertOrUpdate(orderDetail models.OrderDetail) (models.OrderDetail, error) {
	return s.dao.InsertOrUpdate(orderDetail)
}

func (s *orderDetailService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}