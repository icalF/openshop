package usecases

import "github.com/koneko096/openshop/models/datamodels"

type PurchaseValidator interface {
	ValidatePurchase(orderDetail datamodels.OrderDetail) bool
	ValidatePurchaseByProduct(product datamodels.Product, orderDetail datamodels.OrderDetail) bool
}
