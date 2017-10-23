package controllers

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/services"
)

type CouponController struct {
	BaseController
	Service services.CouponService
}

// GET /coupon/
func (c *CouponController) Get() (results []datamodels.Coupon) {
	return c.Service.GetAll()
}
