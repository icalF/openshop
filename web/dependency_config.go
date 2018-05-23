package web

import (
	"go.uber.org/dig"
	"github.com/koneko096/openshop/dao"
	"github.com/koneko096/openshop/datasource"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/koneko096/openshop/web/controllers"
	"github.com/koneko096/openshop/web/middleware"
	"github.com/koneko096/openshop/services"
	"github.com/koneko096/openshop/session"
)

func NewServer(
	userManager services.UserManager,
	couponManager services.CouponManager,
	productManager services.ProductManager,
	paymentManager services.PaymentManager,
	shipmentManager services.ShipmentManager,
	orderDetailManager services.OrderDetailManager,
	orderLalala services.OrderLalala,
	sessionWrapper session.Wrapper,
) *iris.Application {
	app := iris.New()

	//middleware.BasicAuth
	root := app.Party("/")
	root.Controller("/user", new(controllers.UserController),
		userManager,
		sessionWrapper,
	)
	root.Controller("/coupon", new(controllers.CouponController),
		couponManager,
	)
	root.Controller("/product", new(controllers.ProductController),
		productManager,
	)
	root.Controller("/order", new(controllers.OrderController),
		couponManager,
		orderLalala,
		orderDetailManager,
		paymentManager,
		userManager,
		sessionWrapper,
	)
	root.Controller("/shipment", new(controllers.ShipmentController),
		shipmentManager,
	)

	root.Controller("/admin", new(controllers.AdminController),
		orderLalala,
		orderDetailManager,
		paymentManager,
		shipmentManager,
		userManager,
		middleware.BasicAuth,
	)

	app.StaticWeb("/proof", "./uploads")

	return app
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() (*gorm.DB, error){
		dbConn, err := datasource.NewPostgreConnection()
		if err != nil {
			panic(err)
			return nil, err
		}
		return dbConn, nil
	})

	container.Provide(session.NewSessionWrapper)

	container.Provide(dao.NewUserDAO)
	container.Provide(dao.NewCouponDAO)
	container.Provide(dao.NewShipmentDAO)
	container.Provide(dao.NewPaymentDAO)
	container.Provide(dao.NewProductDAO)
	container.Provide(dao.NewOrderDetailDAO)
	container.Provide(dao.NewOrderDAO)

	container.Provide(services.NewUserManager)
	container.Provide(services.NewPaymentProofManager)
	container.Provide(services.NewPaymentManager)
	container.Provide(services.NewCouponValidator)
	container.Provide(services.NewProductManager)
	container.Provide(services.NewCouponManager)
	container.Provide(services.NewShipmentManager)
	container.Provide(services.NewOrderManager)
	container.Provide(services.NewPurchaseValidator)
	container.Provide(services.NewOrderDetailManager)
	container.Provide(services.NewOrderLalala)

	container.Provide(NewServer)

	return container
}
