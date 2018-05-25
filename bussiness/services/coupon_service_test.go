package services

import (
	"testing"
	"time"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/services"
)

func TestCouponService_ValidateCoupon(t *testing.T) {
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

	couponService := services.NewCouponService(nil)
	if couponService.ValidateCoupon(coupons[0]) {
		t.Error("Coupon should be invalid, but considered valid")
	}
	if couponService.ValidateCoupon(coupons[1]) {
		t.Error("Coupon should be invalid, but considered valid")
	}
	if !couponService.ValidateCoupon(coupons[2]) {
		t.Error("Coupon should be valid, but considered invalid")
	}
}
