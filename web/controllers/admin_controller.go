package controllers

import (
	"github.com/kataras/iris"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/services"
)

type AdminController struct {
	BaseController
	OrderService       services.OrderService
	OrderDetailService services.OrderDetailService
	PaymentService     services.PaymentService
	ShipmentService    services.ShipmentService
	UserService        services.UserService
}

// GET /admin/order
func (c *AdminController) GetOrder() (results []datamodels.Order) {
	return c.OrderService.GetAllSubmitted()
}

// GET /admin/order/{id: int}
func (c *AdminController) GetOrderBy(id int64) (interface{}, int) {
	order, found := c.OrderService.GetByID(id)
	if !found {
		return nil, iris.StatusNotFound
	}

	return order, iris.StatusOK
}

// POST /admin/order/{id: int}/cancel
func (c *AdminController) PostOrderByCancel(id int64) (interface{}, int) {
	order, found := c.OrderService.GetByID(id)
	if !found {
		return iris.Map{"message": "order ID not found"}, iris.StatusNotFound
	}

	order.Status = datamodels.CANCELLED
	res, err := c.OrderService.InsertOrUpdate(order)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// GET /admin/order/{id: int}/detail
func (c *AdminController) GetOrderByDetail(orderId int64) (results []datamodels.OrderDetail) {
	return c.OrderDetailService.GetAll()
}

// GET /admin/order/{id: int}/detail/{id: int}
func (c *AdminController) GetOrderByDetailBy(orderId int64, id int64) (interface{}, int) {
	orderDetail, found := c.OrderDetailService.GetByID(id)
	if !found {
		return nil, iris.StatusNotFound
	}

	return orderDetail, iris.StatusOK
}

// GET /admin/order/{id: int}/user
func (c *AdminController) GetOrderByUser(orderId int64) (interface{}, int) {
	order, found := c.OrderService.GetByID(orderId)
	if !found {
		return iris.Map{"message": "order ID not found"}, iris.StatusNotFound
	}

	payment, found := c.UserService.GetByID(order.UserID)
	if !found {
		return iris.Map{"message": "user for order ID not found"}, iris.StatusNotFound
	}

	return payment, iris.StatusOK
}

// GET /admin/order/{id: int}/payment
func (c *AdminController) GetOrderByPayment(orderId int64) (interface{}, int) {
	payment, found := c.PaymentService.GetByOrderID(orderId)
	if !found {
		return iris.Map{"message": "order ID not found"}, iris.StatusNotFound
	}

	return payment, iris.StatusOK
}

// GET /admin/order/{id: int}/payment/proof
func (c *AdminController) PostOrderByPaymentProof(orderId int64) (interface{}, int) {
	payment, found := c.PaymentService.GetByOrderID(orderId)
	if !found {
		return iris.Map{"message": "order ID not found"},
			iris.StatusInternalServerError
	}

	path := c.Ctx.Request().URL.Host + "/proof/" + payment.Proof
	return iris.Map{"path": path}, iris.StatusOK
}

// POST /admin/order/{id: int}/payment/verify
func (c *AdminController) PostOrderByPaymentVerify(orderId int64) (interface{}, int) {
	payment, found := c.PaymentService.GetByOrderID(orderId)
	if !found {
		return iris.Map{"message": "order ID not found"},
			iris.StatusInternalServerError
	}

	payment.Status = datamodels.VERIFIED
	res, err := c.PaymentService.InsertOrUpdate(payment)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// GET /admin/order/{id: int}/shipment
func (c *AdminController) GetOrderByShipment(orderId int64) (interface{}, int) {
	shipment, found := c.ShipmentService.GetByOrderID(orderId)
	if !found {
		return iris.Map{"message": "order ID not found"}, iris.StatusNotFound
	}

	return shipment, iris.StatusOK
}

// POST /admin/order/{id: int}/shipment/send
func (c *AdminController) PostOrderByShipmentSend(orderId int64) (interface{}, int) {
	var shippingCode string
	if err := c.Ctx.ReadJSON(&shippingCode); err != nil || len(shippingCode) != 8 {
		return nil, iris.StatusBadRequest
	}

	shipment := datamodels.NewShipment(orderId, shippingCode)
	res, err := c.ShipmentService.InsertOrUpdate(shipment)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}