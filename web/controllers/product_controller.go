package controllers

import (
	"github.com/kataras/iris"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/bussiness/usecases"
)

type ProductController struct {
	BaseController
	Service usecases.ProductManager
}

// GET /product/
func (c *ProductController) Get() (results []datamodels.Product) {
	return c.Service.GetAllPurchasable()
}

// GET /product/{id: int}
func (c *ProductController) GetBy(id int64) (interface{}, int) {
	product, found := c.Service.GetByID(id)
	if !found {
		return nil, iris.StatusNotFound
	}

	return product, iris.StatusOK
}
