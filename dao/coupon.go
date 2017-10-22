package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type CouponDAO interface {
	Select(query Query) (model models.Coupon, found bool)
	SelectMany(query Query, limit int) (results []models.Coupon)

	InsertOrUpdate(model models.Coupon) (models.Coupon, error)
	Delete(query Query) (deleted bool)
}

type couponRepository struct {
	source *gorm.DB
}

func NewCouponDAO(connection *gorm.DB) CouponDAO {
	return &couponRepository{source: connection}
}

func (r *couponRepository) Select(query Query) (models.Coupon, bool) {
	coupon := models.Coupon{}
	if err := r.source.Where(query).First(&coupon).Error; err != nil {
		return models.Coupon{}, false
	}
	return coupon, true
}

func (r *couponRepository) SelectMany(query Query, limit int) (results []models.Coupon) {
	coupons := new([]models.Coupon)
	r.source.Where(query).Find(&coupons).Limit(limit)
	return *coupons
}

func (r *couponRepository) InsertOrUpdate(coupon models.Coupon) (_ models.Coupon, err error) {
	var oldCoupon models.Coupon
	if err := r.source.First(&oldCoupon).Error; err != nil {
		r.source.Create(&coupon)
	} else {
		r.source.Model(&oldCoupon).Update(&coupon)
	}

	return coupon, err
}

func (r *couponRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.Coupon{}, query).Error; err != nil {
		return false
	}
	return true
}
