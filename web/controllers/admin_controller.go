package controllers

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/services"
	"github.com/kataras/iris"
)

type AdminController struct {
	BaseController
	OrderService services.OrderService
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
		return iris.Map{"message": "order ID not found"},
			iris.StatusInternalServerError
	}

	order.Status = datamodels.CANCELLED
	res, err := c.OrderService.InsertOrUpdate(order)
	if err != nil {
		return iris.Map{"message": err.Error()}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// GET /admin/order/{id: int}/detail
func (c *OrderController) GetOrderByDetail(orderId int64) (results []datamodels.OrderDetail) {
	return c.OrderDetailService.GetAll()
}

// GET /admin/order/{id: int}/detail/{id: int}
func (c *OrderController) GetOrderByDetailBy(orderId int64, id int64) (interface{}, int) {
	orderDetail, found := c.OrderDetailService.GetByID(id)
	if !found {
		return nil, iris.StatusNotFound
	}

	return orderDetail, iris.StatusOK
}
