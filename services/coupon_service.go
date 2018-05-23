package services

import (
	"time"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/dao"
)

func NewCouponManager(dao dao.CouponDAO) CouponManager {
	return &couponService{
		dao: dao,
	}
}

func NewCouponValidator(dao dao.CouponDAO) CouponValidator {
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
