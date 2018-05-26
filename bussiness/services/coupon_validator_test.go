package services

import (
	"testing"
	"time"
	"github.com/koneko096/openshop/models/datamodels"
)

func TestCouponValidator_ValidateCoupon(t *testing.T) {
	coupons := []datamodels.Coupon{
		{
			Qty: 0,
		},
		{
			Qty: 999,
			Due: time.Now().AddDate(0, -1, 0),
		},
		{
			Qty: 1,
			Due: time.Now().AddDate(0, 0, 1),
		},
	}

	couponValidator := NewCouponValidator(nil)
	if couponValidator.ValidateCoupon(coupons[0]) {
		t.Error("Coupon should be invalid, but considered valid")
	}
	if couponValidator.ValidateCoupon(coupons[1]) {
		t.Error("Coupon should be invalid, but considered valid")
	}
	if !couponValidator.ValidateCoupon(coupons[2]) {
		t.Error("Coupon should be valid, but considered invalid")
	}
}
