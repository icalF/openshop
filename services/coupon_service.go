package services

import (
	"time"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type CouponService interface {
	GetAll() []datamodels.Coupon
	GetByID(id int64) (datamodels.Coupon, bool)
	GetByPromoCode(code string) (datamodels.Coupon, bool)
	InsertOrUpdate(coupon datamodels.Coupon) (datamodels.Coupon, error)
	ValidateCoupon(coupon datamodels.Coupon) bool
	ValidateAndTakeCoupon(code string) bool
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

func (s *couponService) InsertOrUpdate(coupon datamodels.Coupon) (datamodels.Coupon, error) {
	return s.dao.InsertOrUpdate(coupon)
}

func (s *couponService) ValidateCoupon(coupon datamodels.Coupon) bool {
	timeNow := time.Now()
	return coupon.Qty > 0 && timeNow.Before(coupon.Due)
}

func (s *couponService) ValidateAndTakeCoupon(code string) bool {
	coupon, _ := s.GetByPromoCode(code)

	if valid := s.ValidateCoupon(coupon); !valid {
		return false
	}

	coupon.Qty -= 1
	if _, err := s.InsertOrUpdate(coupon); err != nil {
		return false
	}

	return true
}

func (s *couponService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}
