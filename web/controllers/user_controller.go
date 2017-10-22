package controllers

import (
	"github.com/kataras/iris"

	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/models/datamodels"
)

type UserController struct {
	BaseController
	Service services.UserService
}

// GET /user/
func (c *UserController) Get() (results []datamodels.User) {
	return c.Service.GetAll()
}

// GET /user/{id: int}
func (c *UserController) GetBy(id int64) (interface{}, int) {
	user, found  := c.Service.GetByID(id);
	if !found {
		return nil, iris.StatusNotFound
	}

	return user, iris.StatusOK
}

// POST /user/
func (c *UserController) Post() (interface{}, int) {
	user := datamodels.User{}
	err := c.Ctx.ReadJSON(&user)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(user)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	res, err := c.Service.InsertOrUpdate(user)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// PUT /user/{id: int}
func (c *UserController) PutBy(id int64) (interface{}, int) {
	user := datamodels.User{}
	err := c.Ctx.ReadJSON(&user)
	if err != nil {
		return "Field(s) parsing error", iris.StatusBadRequest
	}

	err = c.ValidateInput(user)
	if err != nil {
		return err, iris.StatusBadRequest
	}

	user.ID = id
	res, err := c.Service.InsertOrUpdate(user)
	if err != nil {
		return err, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}

// DELETE /user/{id: int}
func (c *UserController) DeleteBy(id int64) (interface{}, int) {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}, iris.StatusAccepted
	}
	return nil, iris.StatusInternalServerError
}