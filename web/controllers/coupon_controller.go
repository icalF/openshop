package controllers

import (
	"github.com/kataras/iris"

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

// GET /coupon/{id: int}
func (c *CouponController) GetBy(id int64) (interface{}, int) {
	coupon, found  := c.Service.GetByID(id);
	if !found {
		return nil, iris.StatusNotFound
	}

	return coupon, iris.StatusOK
}

// POST /coupon/
func (c *CouponController) Post() (interface{}, int) {
	coupon := datamodels.Coupon{}
	err := c.Ctx.ReadJSON(&coupon)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(coupon)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	res, err := c.Service.InsertOrUpdate(coupon)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// PUT /coupon/{id: int}
func (c *CouponController) PutBy(id int64) (interface{}, int) {
	coupon := datamodels.Coupon{}
	err := c.Ctx.ReadJSON(&coupon)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(coupon)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	coupon.ID = id
	res, err := c.Service.InsertOrUpdate(coupon)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /coupon/{id: int}
func (c *CouponController) DeleteBy(id int64) (interface{}, int) {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}, iris.StatusAccepted
	}
	return nil, iris.StatusInternalServerError
}