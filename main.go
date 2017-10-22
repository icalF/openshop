package main

import (
	"log"

	"github.com/kataras/iris"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/icalF/openshop/web/controllers"
	"github.com/icalF/openshop/datasource"
	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/dao"
)

func main() {
	app := iris.New()

	dbConn, err := datasource.NewMysqlConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	userDAO := dao.NewUserDAO(dbConn)
	couponDAO := dao.NewCouponDAO(dbConn)

	userService := services.NewUserService(userDAO)
	couponService := services.NewCouponService(couponDAO)

	app.Controller("/user", new(controllers.UserController),
		userService,
		// Add the basic authentication(admin:password) middleware
		// for the /movies based requests.
		// middleware.BasicAuth
	)

	app.Controller("/coupon", new(controllers.CouponController),
		couponService,
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

	defer dbConn.Close()
}
