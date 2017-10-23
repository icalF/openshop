package controllers

import (
	"github.com/kataras/iris"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/services"
)

type ProductController struct {
	BaseController
	Service services.ProductService
}

// GET /product/
func (c *ProductController) Get() (results []datamodels.Product) {
	return c.Service.GetAllPurchasable()
}

// GET /product/{id: int}
func (c *ProductController) GetBy(id int64) (interface{}, int) {
	product, found  := c.Service.GetByID(id);
	if !found {
		return nil, iris.StatusNotFound
	}

	return product, iris.StatusOK
}

// POST /product/
func (c *ProductController) Post() (interface{}, int) {
	product := datamodels.Product{}
	err := c.Ctx.ReadJSON(&product)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(product)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	res, err := c.Service.InsertOrUpdate(product)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// PUT /product/{id: int}
func (c *ProductController) PutBy(id int64) (interface{}, int) {
	product := datamodels.Product{}
	err := c.Ctx.ReadJSON(&product)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(product)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	product.ID = id
	res, err := c.Service.InsertOrUpdate(product)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /product/{id: int}
func (c *ProductController) DeleteBy(id int64) (interface{}, int) {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}, iris.StatusOK
	}
	return nil, iris.StatusInternalServerError
}