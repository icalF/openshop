package services

import (
	"testing"

	"github.com/koneko096/openshop/services"
	"github.com/koneko096/openshop/models/datamodels"
)

func TestOrderService_GetTotalAmount(t *testing.T) {
	orderDetailService := NewTestOrderDetailService()
	productService := NewTestProductService()
	couponService := NewTestCouponService()
	orderService := services.NewOrderService(nil, nil, orderDetailService, productService, couponService)

	result := orderService.GetTotalAmount(datamodels.Order{ID: 1, VoucherCode:"ABC"})
	if result != 52 {
		t.Errorf("Total amount: expected %d, actual %d", 52, result)
	}
}