package controllers

import (
	"github.com/kataras/iris"
	"github.com/koneko096/openshop/bussiness/usecases"
)

type ShipmentController struct {
	BaseController
	ShipmentService usecases.ShipmentManager
}

// GET /shipment/{code: string}
func (c *ShipmentController) GetBy(code string) (interface{}, int) {
	shipment, found := c.ShipmentService.GetByShippingCode(code)
	if !found {
		return iris.Map{"message": "shipping code not found"}, iris.StatusNotFound
	}

	return shipment, iris.StatusOK
}
