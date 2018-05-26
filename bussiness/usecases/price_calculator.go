package usecases

import "github.com/koneko096/openshop/models/datamodels"

type PriceCalculator interface {
	GetTotalAmount(order datamodels.Order) int
	GetTotalAmountByProductsAndCoupon(
		orderDetails []datamodels.OrderDetail,
		productMap map[int64]datamodels.Product,
		coupon *datamodels.Coupon,
	) int
}
