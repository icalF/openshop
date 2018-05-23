package controllers

import (
	"github.com/koneko096/openshop/services"
	"github.com/kataras/iris"
)

type ShipmentController struct {
	BaseController
	ShipmentService services.ShipmentManager
}

// GET /shipment/{code: string}
func (c *ShipmentController) GetBy(code string) (interface{}, int) {
	shipment, found := c.ShipmentService.GetByShippingCode(code)
	if !found {
		return iris.Map{"message": "shipping code not found"}, iris.StatusNotFound
	}

	return shipment, iris.StatusOK
}
