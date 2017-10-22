package services

import (
	"errors"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type OrderService interface {
	GetAll() []datamodels.Order
	GetByID(id int64) (datamodels.Order, bool)
	InsertOrUpdate(order datamodels.Order) (datamodels.Order, error)
	InsertCoupon(id int64, code string) (bool, error)
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

func (s *orderService) InsertOrUpdate(order datamodels.Order) (datamodels.Order, error) {
	return s.dao.InsertOrUpdate(order)
}

func (s *orderService) InsertCoupon(id int64, code string) (bool, error) {
	order, found := s.GetByID(id)
	if !found {
		return false, errors.New("coupon code couldn't be found")
	}
	if order.Status != datamodels.UNSUBMITTED {
		return false, errors.New("coupon cannot be applied to submitted order")
	}

	order.VoucherCode = code
	_, err := s.dao.InsertOrUpdate(order)
	return true, err
}

func (s *orderService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}
