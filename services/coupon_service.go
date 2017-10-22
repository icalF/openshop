package services

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type CouponService interface {
	GetAll() []datamodels.Coupon
	GetByID(id int64) (datamodels.Coupon, bool)
	GetByPromoCode(code string) (datamodels.Coupon, bool)
	GetByOrderID(orderID int64) (datamodels.Coupon, bool)
	InsertOrUpdate(coupon datamodels.Coupon) (datamodels.Coupon, error)
	DeleteByID(id int64) bool
}

func NewCouponService(dao dao.CouponDAO, orderService OrderService) CouponService {
	return &couponService{
		dao: dao,
		orderService: orderService,
	}
}

type couponService struct {
	dao dao.CouponDAO
	orderService OrderService
}

func (s *couponService) GetAll() []datamodels.Coupon {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *couponService) GetByID(id int64) (datamodels.Coupon, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *couponService) GetByPromoCode(code string) (datamodels.Coupon, bool) {
	return s.dao.Select(map[string]string{
		"code": code,
	})
}

func (s *couponService) GetByOrderID(orderID int64) (datamodels.Coupon, bool) {
	order, found := s.orderService.GetByID(orderID)
	if !found {
		return datamodels.Coupon{}, false
	}

	return s.GetByPromoCode(order.VoucherCode)
}

func (s *couponService) InsertOrUpdate(coupon datamodels.Coupon) (datamodels.Coupon, error) {
	return s.dao.InsertOrUpdate(coupon)
}

func (s *couponService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}
