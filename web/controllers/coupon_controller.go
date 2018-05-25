package controllers

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/bussiness/usecases"
)

type CouponController struct {
	BaseController
	Service usecases.CouponManager
}

// GET /coupon/
func (c *CouponController) Get() (results []datamodels.Coupon) {
	return c.Service.GetAll()
}
