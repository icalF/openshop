package services

import (
	"errors"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/dao"
	"github.com/koneko096/openshop/bussiness/usecases"
)


func NewOrderManager(dao dao.OrderDAO) usecases.OrderManager {
	return &orderService{
		dao: dao,
	}
}

func NewOrderLalala(dao dao.OrderDAO,
	paymentManager usecases.PaymentManager,
	orderDetailManager usecases.OrderDetailManager,
	productManager usecases.ProductManager,
	couponManager usecases.CouponManager,
	couponValidator usecases.CouponValidator,
	purchaseValidator usecases.PurchaseValidator,
) usecases.OrderLalala {
	return &orderService{
		dao:                dao,
		paymentManager:     paymentManager,
		orderDetailManager: orderDetailManager,
		productManager:     productManager,
		couponManager:      couponManager,
		couponValidator:    couponValidator,
		purchaseValidator:  purchaseValidator,
	}
}


type orderService struct {
	dao                dao.OrderDAO
	paymentManager     usecases.PaymentManager
	orderDetailManager usecases.OrderDetailManager
	productManager     usecases.ProductManager
	couponManager      usecases.CouponManager
	couponValidator    usecases.CouponValidator
	purchaseValidator  usecases.PurchaseValidator
}



func (s *orderService) GetAll() []datamodels.Order {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *orderService) GetAllSubmitted() []datamodels.Order {
	return s.dao.SelectMany(map[string]string{
		"status": datamodels.SUBMITTED,
	}, 0)
}

func (s *orderService) GetByID(id int64) (datamodels.Order, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *orderService) GetByUserID(userId int64) []datamodels.Order {
	return s.dao.SelectMany(map[string]string{
		"user_id": string(userId),
	}, 0)
}

func (s *orderService) InsertOrUpdate(order datamodels.Order) (datamodels.Order, error) {
	return s.dao.InsertOrUpdate(order)
}

func (s *orderService) InsertCoupon(id int64, code string) (bool, error) {
	order, found := s.GetByID(id)
	if !found {
		return false, errors.New("order with ID couldn't be found")
	}

	if order.Status != datamodels.UNSUBMITTED {
		return false, errors.New("coupon cannot be applied to submitted order")
	}

	coupon, found := s.couponManager.GetByPromoCode(order.VoucherCode)
	if !found {
		return false, errors.New("promo code not found")
	}

	if valid := s.couponValidator.ValidateCoupon(coupon); !valid {
		return true, errors.New("promo code does not valid")
	}

	order.VoucherCode = code
	_, err := s.dao.InsertOrUpdate(order)
	return true, err
}

func (s *orderService) Checkout(id int64) (datamodels.Payment, bool, error) {
	order, found := s.GetByID(id)
	if !found {
		return datamodels.Payment{}, false, errors.New("order with ID couldn't be found")
	}
	if order.Status != datamodels.UNSUBMITTED {
		return datamodels.Payment{}, false, errors.New("checkout cannot be done to submitted order")
	}

	if success := s.couponValidator.ValidateAndTakeCoupon(order.VoucherCode); !success {
		return datamodels.Payment{}, true, errors.New("coupon code does not valid")
	}
	if success := s.ValidateOrderDetails(id); !success {
		return datamodels.Payment{}, true, errors.New("product(s) cannot be purchased")
	}

	payment, err := s.paymentManager.InsertOrUpdate(datamodels.NewPayment(id, s.GetTotalAmount(order)))
	if err != nil {
		return datamodels.Payment{}, true, err
	}

	order.Status = datamodels.SUBMITTED
	if _, err = s.dao.InsertOrUpdate(order); err != nil {
		return datamodels.Payment{}, true, err
	}

	return payment, true, nil
}

func (s *orderService) ValidateOrderDetails(orderId int64) bool {
	orderDetails := s.orderDetailManager.GetByOrderID(orderId)
	for _, orderDetail := range orderDetails {
		if !s.purchaseValidator.ValidatePurchase(orderDetail) {
			return false
		}
	}
	return true
}

func (s *orderService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}

func (s *orderService) GetTotalAmount(order datamodels.Order) int {
	orderDetails := s.orderDetailManager.GetByOrderID(order.ID)
	productIds := make([]int64, len(orderDetails))
	for i, orderDetail := range orderDetails {
		productIds[i] = orderDetail.ProductID
	}

	products := s.productManager.GetByIDs(productIds)
	productMap := s.productManager.CreateProductMap(products)

	grossAmount := 0
	for _, orderDetail := range orderDetails {
		grossAmount += productMap[orderDetail.ProductID].Price * orderDetail.Qty
	}

	coupon, found := s.couponManager.GetByPromoCode(order.VoucherCode)
	if found {
		if coupon.Percent > 0 {
			grossAmount = int(grossAmount * (100 - coupon.Percent) / 100.0)
		} else {
			grossAmount -= coupon.Nominal
		}
	}
	return grossAmount
}
