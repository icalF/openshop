package services

import (
	"testing"

	"github.com/koneko096/openshop/services"
	"github.com/koneko096/openshop/models/datamodels"
)

var products = []datamodels.Product{
	{
		ID:  0,
		Qty: 0,
	},
	{
		ID:  1,
		Qty: 1,
	},
	{
		ID:  2,
		Qty: 2,
	},
	{
		ID:  3,
		Qty: 3,
	},
}

func TestProductService_CreateProductMap(t *testing.T) {
	service := services.NewProductService(nil)

	productMap := service.CreateProductMap(products)
	for i := 0; i <= 3; i++ {
		if productMap[int64(i)].Qty != i {
			t.Errorf("Element ID(%d) qty: expected %d, actual %d", i, i, productMap[int64(i)].Qty)
		}
	}
}
