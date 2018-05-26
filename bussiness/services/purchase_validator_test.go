package services

import (
	"testing"
	"github.com/koneko096/openshop/models/datamodels"
)

func TestPurchaseValidator_ValidatePurchase(t *testing.T) {
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

	products := []datamodels.Product{
		{
			ID:    1,
			Price: 35,
			Qty:   3,
		},
		{
			ID:    2,
			Price: 15,
			Qty:   1,
		},
		{
			ID:    3,
			Price: 5,
		},
	}

	purchaseValidator := NewPurchaseValidator(nil, nil)
	if !purchaseValidator.ValidatePurchaseByProduct(products[0], orderDetails[0]) {
		t.Error("Coupon should be valid, but considered invalid")
	}
	if purchaseValidator.ValidatePurchaseByProduct(products[1], orderDetails[1]) {
		t.Error("Coupon should be invalid, but considered valid")
	}
}
