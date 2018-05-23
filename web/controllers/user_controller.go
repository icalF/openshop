package controllers

import (
	"github.com/kataras/iris"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/services"
	"github.com/koneko096/openshop/session"
)

type UserController struct {
	BaseController
	UserService    services.UserManager
	SessionWrapper session.Wrapper
}

// GET /user
func (c *UserController) Get() (interface{}, int) {
	sess := c.SessionWrapper.GetSession().Start(c.Ctx)

	user, err := c.UserService.GetByToken(sess.ID())
	if err != nil {
		return iris.Map{"message": "unexpected error"}, iris.StatusInternalServerError
	}

	return user, iris.StatusOK
}

// PUT /user
func (c *UserController) Put() (interface{}, int) {
	sess := c.SessionWrapper.GetSession().Start(c.Ctx)

	user := datamodels.User{}
	err := c.Ctx.ReadJSON(&user)
	if err != nil {
		return iris.Map{"message": "field(s) parsing error"}, iris.StatusBadRequest
	}

	err = c.ValidateInput(user)
	if err != nil {
		return iris.Map{"message": err.Error}, iris.StatusBadRequest
	}

	oldUser, err := c.UserService.GetByToken(sess.ID())
	if err != nil {
		return iris.Map{"message": "unexpected error"}, iris.StatusInternalServerError
	}

	user.ID = oldUser.ID
	res, err := c.UserService.InsertOrUpdate(user)
	if err != nil {
		return iris.Map{"message": err.Error}, iris.StatusInternalServerError
	}

	return res, iris.StatusOK
}
