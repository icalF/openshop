package services

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type OrderService interface {
	GetAll() []datamodels.Order
	GetByID(id int64) (datamodels.Order, bool)
	Insert(order datamodels.Order) (datamodels.Order, error)
	DeleteByID(id int64) bool
}

func NewOrderService(dao dao.OrderDAO) OrderService {
	return &orderService{
		dao: dao,
	}
}

type orderService struct {
	dao dao.OrderDAO
}

func (s *orderService) GetAll() []datamodels.Order {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *orderService) GetByID(id int64) (datamodels.Order, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *orderService) Insert(order datamodels.Order) (datamodels.Order, error) {
	return s.dao.Insert(order)
}

func (s *orderService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}