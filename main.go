package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/koneko096/openshop/web"
	"github.com/kataras/iris"
)

func main() {
	//defer dbConn.Close()

	container := web.BuildContainer()

	err := container.Invoke(func(app *iris.Application) {
		app.Run(
			iris.Addr(":"+os.Getenv("PORT")),
			iris.WithoutVersionChecker,
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
		)
	})

	if err != nil {
		panic(err)
	}
}
