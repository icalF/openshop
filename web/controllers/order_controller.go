package controllers

import (
	"github.com/kataras/iris"
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/bussiness/usecases"
	"github.com/koneko096/openshop/web/session"
)

type OrderController struct {
	BaseController
	CouponService       usecases.CouponManager
	OrderService        usecases.OrderManager
	OrderLalala         usecases.OrderLalala
	OrderDetailService  usecases.OrderDetailManager
	PurchaseValidator   usecases.PurchaseValidator
	CouponValidator     usecases.CouponValidator
	PaymentService      usecases.PaymentManager
	PaymentProofService usecases.PaymentProofManager
	UserService         usecases.UserManager
	SessionWrapper      session.Wrapper
}

// GET /order
func (c *OrderController) Get() (interface{}, int) {
	sess := c.SessionWrapper.GetSession().Start(c.Ctx)

	user, err := c.UserService.GetByToken(sess.ID())
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return c.OrderService.GetByUserID(user.ID), iris.StatusOK
}

// POST /order
func (c *OrderController) Post() (interface{}, int) {
	sess := c.SessionWrapper.GetSession().Start(c.Ctx)

	order := datamodels.Order{}
	err := c.Ctx.ReadJSON(&order)
	if err != nil {
		return iris.Map{"message": "field(s) parsing error"}, iris.StatusBadRequest
	}

	err = c.ValidateInput(order)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusBadRequest
	}

	user, err := c.UserService.GetByToken(sess.ID())
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	res, err := c.OrderService.InsertOrUpdate(datamodels.NewOrder(user.ID))
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /order/{id: int}
func (c *OrderController) DeleteBy(id int64) (interface{}, int) {
	wasDel := c.OrderService.DeleteByID(id)
	if wasDel {
		return iris.Map{"order_id": id}, iris.StatusOK
	}
	return nil, iris.StatusInternalServerError
}

// GET /order/{id: int}/coupon
func (c *OrderController) GetByCoupon(orderId int64) (interface{}, int) {
	order, found := c.OrderService.GetByID(orderId)
	if !found {
		return iris.Map{"message": "order ID not found"}, iris.StatusNotFound
	}

	coupon, found := c.CouponService.GetByPromoCode(order.VoucherCode)
	if !found {
		return iris.Map{"message": "promo code not found"}, iris.StatusNotFound
	}

	return coupon, iris.StatusOK
}

// POST /order/{id: int}/coupon
func (c *OrderController) PostByCoupon(orderId int64) (interface{}, int) {
	var code string
	if err := c.Ctx.ReadJSON(&code); err != nil || len(code) != 6 {
		return nil, iris.StatusBadRequest
	}

	found, err := c.OrderLalala.InsertCoupon(orderId, code)
	if !found {
		return iris.Map{"message": err.Error()}, iris.StatusNotFound
	}
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return iris.Map{"voucher_code": code}, iris.StatusOK
}

// DELETE /order/{id: int}/coupon
func (c *OrderController) DeleteByCoupon(orderId int64) (interface{}, int) {
	order, found := c.OrderService.GetByID(orderId)
	if !found {
		return nil, iris.StatusNotFound
	}

	oldCode := order.VoucherCode
	order.VoucherCode = ""
	if _, err := c.OrderService.InsertOrUpdate(order); err != nil {
		return nil, iris.StatusInternalServerError
	}

	return iris.Map{"voucher_code": oldCode}, iris.StatusOK
}

// POST /order/{id: int}/checkout
func (c *OrderController) PostByCheckout(orderId int64) (interface{}, int) {
	payment, valid, err := c.OrderLalala.Checkout(orderId)
	if !valid {
		return iris.Map{"message": err.Error()}, iris.StatusBadRequest
	}
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return payment, iris.StatusOK
}

// GET /order/{id: int}/payment
func (c *OrderController) GetByPayment(orderId int64) (interface{}, int) {
	payment, found := c.PaymentService.GetByOrderID(orderId)
	if !found {
		return iris.Map{"message": "orderId not found"}, iris.StatusNotFound
	}

	return payment, iris.StatusOK
}

// POST /order/{id: int}/payment/upload
func (c *OrderController) PostByPaymentUpload(orderId int64) (interface{}, int) {
	c.Ctx.SetMaxRequestBodySize(1 << 16)

	// Get the file from the request
	file, _, err := c.Ctx.FormFile("file")
	if err != nil {
		return iris.Map{
			"message": "error while uploading",
			"info":    err.Error(),
		}, iris.StatusInternalServerError
	}
	defer file.Close()

	found, err := c.PaymentProofService.UpdatePaymentProof(orderId, file)
	if err != nil {
		statusCode := iris.StatusInternalServerError
		if !found {
			statusCode = iris.StatusNotFound
		}

		return iris.Map{
			"message": "error while updating payment data",
			"info":    err.Error(),
		}, statusCode
	}

	return iris.Map{"message": "uploading success"}, iris.StatusOK
}

// GET /order/{id: int}/detail
func (c *OrderController) GetByDetail(orderId int64) (results []datamodels.OrderDetail) {
	return c.OrderDetailService.GetAll()
}

// POST /order/{id: int}/detail
func (c *OrderController) PostByDetail(orderId int64) (interface{}, int) {
	orderDetail := datamodels.NewOrderDetail(orderId)
	err := c.Ctx.ReadJSON(&orderDetail)
	if err != nil {
		return iris.Map{"message": "field(s) parsing error"}, iris.StatusBadRequest
	}

	err = c.ValidateInput(orderDetail)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusBadRequest
	}

	if valid := c.PurchaseValidator.ValidatePurchase(orderDetail); !valid {
		return iris.Map{"message": "input not valid"}, iris.StatusBadRequest
	}

	res, err := c.OrderDetailService.InsertOrUpdate(orderDetail)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// GET /order/{id: int}/detail/{id: int}
func (c *OrderController) GetByDetailBy(orderId int64, id int64) (interface{}, int) {
	orderDetail, found := c.OrderDetailService.GetByID(id)
	if !found {
		return iris.Map{"message": "order detail ID cannot be found"}, iris.StatusNotFound
	}

	return orderDetail, iris.StatusOK
}

// PUT /order/{id: int}/detail/{id: int}
func (c *OrderController) PutByDetailBy(orderId int64, id int64) (interface{}, int) {
	orderDetail := datamodels.NewOrderDetail(orderId)
	err := c.Ctx.ReadJSON(&orderDetail)
	if err != nil {
		return iris.Map{"message": "field(s) parsing error"}, iris.StatusBadRequest
	}

	err = c.ValidateInput(orderDetail)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusBadRequest
	}

	if valid := c.PurchaseValidator.ValidatePurchase(orderDetail); !valid {
		return iris.Map{"message": "input not valid"}, iris.StatusBadRequest
	}

	orderDetail.ID = id
	res, err := c.OrderDetailService.InsertOrUpdate(orderDetail)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /order/{id: int}/detail/{id: int}
func (c *OrderController) DeleteByDetailBy(orderId int64, id int64) (interface{}, int) {
	wasDel := c.OrderDetailService.DeleteByID(id)
	if wasDel {
		return iris.Map{"order_detail_id": id}, iris.StatusOK
	}
	return nil, iris.StatusInternalServerError
}
