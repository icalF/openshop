package middleware

import "github.com/kataras/iris/middleware/basicauth"

var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"user":  "customer123",
		"admin": "admin123",
	},
})
