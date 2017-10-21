package main

import (
	"github.com/kataras/iris"

	"github.com/icalF/openshop/web/controllers"
	"github.com/icalF/openshop/datasource"
	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/dao"
)

func main() {
	app := iris.New()

	userDAO := dao.NewUserDAO(datasource.Users)
	userService := services.NewUserService(userDAO)

	app.Controller("/user", new(controllers.UserController),
		userService,
		// Add the basic authentication(admin:password) middleware
		// for the /movies based requests.
		// middleware.BasicAuth
		)
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)

}
