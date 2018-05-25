package usecases

import "github.com/koneko096/openshop/models/datamodels"

type CouponManager interface {
	GetAll() []datamodels.Coupon
	GetByID(id int64) (datamodels.Coupon, bool)
	GetByPromoCode(code string) (datamodels.Coupon, bool)
	InsertOrUpdate(coupon datamodels.Coupon) (datamodels.Coupon, error)
	DeleteByID(id int64) bool
}
