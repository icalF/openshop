package services

import (
	"errors"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type OrderService interface {
	GetAll() []datamodels.Order
	GetAllSubmitted() []datamodels.Order
	GetByID(id int64) (datamodels.Order, bool)
	GetByUserID(userId int64) []datamodels.Order
	InsertOrUpdate(order datamodels.Order) (datamodels.Order, error)
	InsertCoupon(id int64, code string) (bool, error)
	Checkout(id int64) (datamodels.Payment, bool, error)
	DeleteByID(id int64) bool
	GetTotalAmount(order datamodels.Order) int
}

func NewOrderService(dao dao.OrderDAO,
	paymentService PaymentService,
	orderDetailService OrderDetailService,
	productService ProductService,
	couponService CouponService,
) OrderService {
	return &orderService{
		dao:                dao,
		paymentService:     paymentService,
		orderDetailService: orderDetailService,
		productService:     productService,
		couponService:      couponService,
	}
}

type orderService struct {
	dao                dao.OrderDAO
	paymentService     PaymentService
	orderDetailService OrderDetailService
	productService     ProductService
	couponService      CouponService
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

	coupon, found := s.couponService.GetByPromoCode(order.VoucherCode)
	if !found {
		return false, errors.New("promo code not found")
	}

	if valid := s.couponService.ValidateCoupon(coupon); !valid {
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

	if success := s.couponService.ValidateAndTakeCoupon(order.VoucherCode); !success {
		return datamodels.Payment{}, true, errors.New("coupon code does not valid")
	}
	if success := s.ValidateOrderDetails(id); !success {
		return datamodels.Payment{}, true, errors.New("product(s) cannot be purchased")
	}

	payment, err := s.paymentService.InsertOrUpdate(datamodels.NewPayment(id, s.GetTotalAmount(order)))
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
	orderDetails := s.orderDetailService.GetByOrderID(orderId)
	for _, orderDetail := range orderDetails {
		if !s.orderDetailService.ValidatePurchase(orderDetail) {
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
	orderDetails := s.orderDetailService.GetByOrderID(order.ID)
	productIds := make([]int64, len(orderDetails))
	for i, orderDetail := range orderDetails {
		productIds[i] = orderDetail.ProductID
	}

	products := s.productService.GetByIDs(productIds)
	productMap := s.productService.CreateProductMap(products)

	grossAmount := 0
	for _, orderDetail := range orderDetails {
		grossAmount += productMap[orderDetail.ProductID].Price * orderDetail.Qty
	}

	coupon, found := s.couponService.GetByPromoCode(order.VoucherCode)
	if found {
		if coupon.Percent > 0 {
			grossAmount = int(grossAmount * (100 - coupon.Percent) / 100.0)
		} else {
			grossAmount -= coupon.Nominal
		}
	}
	return grossAmount
}
