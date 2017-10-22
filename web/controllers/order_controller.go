package controllers

import (
	"github.com/kataras/iris"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/services"
)

type OrderController struct {
	BaseController
	CouponService services.CouponService
	OrderService services.OrderService
	OrderDetailService services.OrderDetailService
	PaymentService services.PaymentService
}

// GET /order/
func (c *OrderController) Get() (results []datamodels.Order) {
	return c.OrderService.GetAll()
}

// GET /order/{id: int}
func (c *OrderController) GetBy(id int64) (interface{}, int) {
	order, found := c.OrderService.GetByID(id)
	if !found {
		return nil, iris.StatusNotFound
	}

	return order, iris.StatusOK
}

// POST /order/
func (c *OrderController) Post() (interface{}, int) {
	order := datamodels.Order{}
	err := c.Ctx.ReadJSON(&order)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(order)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	res, err := c.OrderService.InsertOrUpdate(datamodels.NewOrder(order.UserID))
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /order/{id: int}
func (c *OrderController) DeleteBy(id int64) (interface{}, int) {
	wasDel := c.OrderService.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}, iris.StatusOK
	}
	return nil, iris.StatusInternalServerError
}

// GET /order/{id: int}/detail/
func (c *OrderController) GetByDetail(orderId int64) (results []datamodels.OrderDetail) {
	return c.OrderDetailService.GetAll()
}

// POST /order/{id: int}/detail/
func (c *OrderController) PostByDetail(orderId int64) (interface{}, int) {
	orderDetail := datamodels.NewOrderDetail(orderId)
	err := c.Ctx.ReadJSON(&orderDetail)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(orderDetail)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	res, err := c.OrderDetailService.InsertOrUpdate(orderDetail)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// GET /order/{id: int}/coupon
func (c *OrderController) GetByCoupon(orderId int64) (interface{}, int) {
	coupon, found := c.CouponService.GetByOrderID(orderId)
	if !found {
		return nil, iris.StatusNotFound
	}

	return coupon, iris.StatusOK
}

// POST /order/{id: int}/coupon
func (c *OrderController) PostByCoupon(orderId int64) (interface{}, int) {
	var code string
	if err := c.Ctx.ReadJSON(&code); err != nil || len(code) < 1 {
		return nil, iris.StatusBadRequest
	}

	found, err := c.OrderService.InsertCoupon(orderId, code)
	if !found {
		return iris.Map{"message": "coupon with code not found"}, iris.StatusNotFound
	}
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return iris.Map{"couponApplied": code}, iris.StatusOK
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

	return iris.Map{"deleted": oldCode}, iris.StatusOK
}

// GET /order/{id: int}/detail/{id: int}
func (c *OrderController) GetByDetailBy(orderId int64, id int64) (interface{}, int) {
	orderDetail, found := c.OrderDetailService.GetByID(id)
	if !found {
		return nil, iris.StatusNotFound
	}

	return orderDetail, iris.StatusOK
}

// PUT /order/{id: int}/detail/{id: int}
func (c *OrderController) PutByDetailBy(orderId int64, id int64) (interface{}, int) {
	orderDetail := datamodels.NewOrderDetail(orderId)
	err := c.Ctx.ReadJSON(&orderDetail)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(orderDetail)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	orderDetail.ID = id
	res, err := c.OrderDetailService.InsertOrUpdate(orderDetail)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /order/{id: int}/detail/{id: int}
func (c *OrderController) DeleteByDetailBy(orderId int64, id int64) (interface{}, int) {
	wasDel := c.OrderDetailService.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}, iris.StatusOK
	}
	return nil, iris.StatusInternalServerError
}