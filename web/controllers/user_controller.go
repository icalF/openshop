package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/models/datamodels"
)

type UserController struct {
	BaseController
	sess        *sessions.Sessions
	UserService services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		sess: sessions.New(sessions.Config{Cookie: "SHOPSESS_ID"}),
	}
}

// GET /user
func (c *UserController) Get() (interface{}, int) {
	sess := c.sess.Start(c.Ctx)

	user, err := c.UserService.GetByToken(sess.ID())
	if err != nil {
		return iris.Map{"message": "unexpected error"}, iris.StatusInternalServerError
	}

	return user, iris.StatusOK
}

// PUT /user
func (c *UserController) Post() (interface{}, int) {
	sess := c.sess.Start(c.Ctx)

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