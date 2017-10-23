package controllers

import (
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v8"
)

type BaseController struct {
	mvc.C
}

func (c *BaseController) ValidateInput(model interface{}) error {
	validate := validator.New()
	validate.SetTagName("validate")

	return validate.Struct(model)
}
