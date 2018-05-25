package usecases

import "github.com/koneko096/openshop/models/datamodels"

type CouponValidator interface {
	ValidateCoupon(coupon datamodels.Coupon) bool
	ValidateAndTakeCoupon(code string) bool
}
