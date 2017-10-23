package main

import (
	"log"

	"github.com/kataras/iris"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/icalF/openshop/web/controllers"
	"github.com/icalF/openshop/datasource"
	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/dao"
	"github.com/icalF/openshop/web/middleware"
)

func main() {
	app := iris.New()

	dbConn, err := datasource.NewPostgreConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	userDAO := dao.NewUserDAO(dbConn)
	couponDAO := dao.NewCouponDAO(dbConn)
	productDAO := dao.NewProductDAO(dbConn)
	shipmentDAO := dao.NewShipmentDAO(dbConn)
	paymentDAO := dao.NewPaymentDAO(dbConn)
	orderDetailDAO := dao.NewOrderDetailDAO(dbConn)
	orderDAO := dao.NewOrderDAO(dbConn)

	userService := services.NewUserService(userDAO)
	shipmentService := services.NewShipmentService(shipmentDAO)
	productService := services.NewProductService(productDAO)
	paymentService := services.NewPaymentService(paymentDAO)
	orderDetailService := services.NewOrderDetailService(orderDetailDAO, productService)
	couponService := services.NewCouponService(couponDAO)
	orderService := services.NewOrderService(orderDAO, paymentService, orderDetailService, productService, couponService)

	//middleware.BasicAuth
	root := app.Party("/")
	root.Controller("/user", controllers.NewUserController(),
		userService,
	)
	root.Controller("/coupon", new(controllers.CouponController),
		couponService,
	)
	root.Controller("/product", new(controllers.ProductController),
		productService,
	)
	root.Controller("/order", controllers.NewOrderController(),
		couponService,
		orderService,
		orderDetailService,
		paymentService,
		userService,
	)
	root.Controller("/shipment", new(controllers.ShipmentController),
		shipmentService,
	)

	root.Controller("/admin", new(controllers.AdminController),
		orderService,
		orderDetailService,
		paymentService,
		shipmentService,
		userService,
		middleware.BasicAuth,
	)

	app.StaticWeb("/proof", "./uploads")

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

	defer dbConn.Close()
}
