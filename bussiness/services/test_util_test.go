package services

import (
	"github.com/koneko096/openshop/models/datamodels"
)

func CreateProductMap(products []datamodels.Product) map[int64]datamodels.Product {
	productMap := map[int64]datamodels.Product{}
	for _, product := range products {
		productMap[product.ID] = product
	}
	return productMap
}