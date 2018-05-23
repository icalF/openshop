package services_test

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/services"
	"testing"
)

func TestOrderDetailService_ValidatePurchase(t *testing.T) {
	orderDetails := []datamodels.OrderDetail{
		{
			ProductID: 1,
			Qty:       2,
		},
		{
			ProductID: 2,
			Qty:       2,
		},
	}

	productService := NewTestProductService()
	orderDetailService := services.NewOrderDetailService(nil, productService)
	if !orderDetailService.ValidatePurchase(orderDetails[0]) {
		t.Error("Coupon should be valid, but considered invalid")
	}
	if orderDetailService.ValidatePurchase(orderDetails[1]) {
		t.Error("Coupon should be invalid, but considered valid")
	}
}
