package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"

	"github.com/icalF/openshop/dao"
	"github.com/icalF/openshop/datasource"
	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/web/controllers"
	"github.com/icalF/openshop/web/middleware"
	"github.com/icalF/openshop/session"
)

func main() {
	app := iris.New()

	dbConn, err := datasource.NewPostgreConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	sessionWrapper := session.NewSessionWrapper()

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
	root.Controller("/user", new(controllers.UserController),
		userService,
		sessionWrapper,
	)
	root.Controller("/coupon", new(controllers.CouponController),
		couponService,
	)
	root.Controller("/product", new(controllers.ProductController),
		productService,
	)
	root.Controller("/order", new(controllers.OrderController),
		couponService,
		orderService,
		orderDetailService,
		paymentService,
		userService,
		sessionWrapper,
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
		iris.Addr(":"+os.Getenv("PORT")),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

	defer dbConn.Close()
}
