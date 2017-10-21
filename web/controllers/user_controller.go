package controllers

import (
	"errors"
	_ "github.com/gpmgo/gopm/modules/log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/services"
)

type UserController struct {
	mvc.C
	Service services.UserService
}

// GET /user/
func (c *UserController) Get() (results []models.User) {
	return c.Service.GetAll()
}

// GET /user/{id: int}
func (c *UserController) GetBy(id int64) (movie models.User, found bool) {
	return c.Service.GetByID(id)
}

// POST /user/
func (c *UserController) Post() (models.User, error) {
	user := models.User{}
	err := c.Ctx.ReadJSON(&user)
	if err != nil {
		return models.User{}, errors.New("field(s) parsing error")
	}

	return c.Service.InsertOrUpdate(user)
}

// PUT /user/{id: int}
func (c *UserController) PutBy(id int64) (models.User, error) {
	user := models.User{}
	err := c.Ctx.ReadJSON(&user)
	if err != nil {
		return models.User{}, errors.New("field(s) parsing error")
	}

	user.ID = id
	return c.Service.InsertOrUpdate(user)
}

// DELETE /user/{id: int}
func (c *UserController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}
	}
	return iris.StatusBadRequest
}