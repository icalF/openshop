package controllers

import (
	"gopkg.in/go-playground/validator.v8"
	"github.com/kataras/iris/mvc"
)

type BaseController struct {
	mvc.C
}

func (c *BaseController) ValidateInput(model interface{}) (error) {
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)

	return validate.Struct(model)
}