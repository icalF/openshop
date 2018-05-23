package dao

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type CouponDAO interface {
	Select(query Query) (model datamodels.Coupon, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Coupon)

	InsertOrUpdate(model datamodels.Coupon) (datamodels.Coupon, error)
	Delete(query Query) (deleted bool)
}

type couponRepository struct {
	source *gorm.DB
}

func NewCouponDAO(connection *gorm.DB) CouponDAO {
	return &couponRepository{source: connection}
}

func (r *couponRepository) Select(query Query) (datamodels.Coupon, bool) {
	coupon := datamodels.Coupon{}
	if err := r.source.Where(query).First(&coupon).Error; err != nil {
		return datamodels.Coupon{}, false
	}
	return coupon, true
}

func (r *couponRepository) SelectMany(query Query, limit int) (results []datamodels.Coupon) {
	coupons := new([]datamodels.Coupon)
	r.source.Where(query).Find(&coupons).Limit(limit)
	return *coupons
}

func (r *couponRepository) InsertOrUpdate(coupon datamodels.Coupon) (_ datamodels.Coupon, err error) {
	var oldCoupon datamodels.Coupon
	if err := r.source.First(&oldCoupon).Error; err != nil {
		r.source.Create(&coupon)
	} else {
		r.source.Model(&oldCoupon).Update(&coupon)
	}

	return coupon, err
}

func (r *couponRepository) Delete(query Query) bool {
	if err := r.source.Delete(datamodels.Coupon{}, query).Error; err != nil {
		return false
	}
	return true
}
