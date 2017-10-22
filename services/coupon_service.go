package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type CouponService interface {
	GetAll() []models.Coupon
	GetByID(id int64) (models.Coupon, bool)
	InsertOrUpdate(coupon models.Coupon) (models.Coupon, error)
	DeleteByID(id int64) bool
}

func NewCouponService(dao dao.CouponDAO) CouponService {
	return &couponService{
		dao: dao,
	}
}

type couponService struct {
	dao dao.CouponDAO
}

func (s *couponService) GetAll() []models.Coupon {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *couponService) GetByID(id int64) (models.Coupon, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *couponService) InsertOrUpdate(coupon models.Coupon) (models.Coupon, error) {
	return s.dao.InsertOrUpdate(coupon)
}

func (s *couponService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}