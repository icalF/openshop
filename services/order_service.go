package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type OrderService interface {
	GetAll() []models.Order
	GetByID(id int64) (models.Order, bool)
	Insert(order models.Order) (models.Order, error)
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

func (s *orderService) GetAll() []models.Order {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *orderService) GetByID(id int64) (models.Order, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *orderService) Insert(order models.Order) (models.Order, error) {
	return s.dao.Insert(order)
}

func (s *orderService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}