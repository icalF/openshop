package services

import (
	"testing"
	"github.com/koneko096/openshop/models/datamodels"
)

func TestPriceCalculator_GetTotalAmount(t *testing.T) {
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

	orderDetails := []datamodels.OrderDetail{
		{
			ProductID: 1,
			Qty:       2,
		},
		{
			ProductID: 2,
			Qty:       2,
		},
		{
			ProductID: 3,
			Qty:       1,
		},
	}

	promoCodes := map[string]datamodels.Coupon{
		"ABC": {
			Percent: 50,
		},
		"DEF": {
			Nominal: 20000,
		},
	}
	coupon := promoCodes["ABC"]

	priceCalculator := NewPriceCalculator(nil, nil, nil)

	result := priceCalculator.GetTotalAmountByProductsAndCoupon(orderDetails, CreateProductMap(products), &coupon)
	if result != 52 {
		t.Errorf("Total amount: expected %d, actual %d", 52, result)
	}
}