package main

import (
	"os"

	"github.com/koneko096/openshop/web/config"
	"github.com/kataras/iris"
)

func main() {
	container := config.BuildContainer()

	port := "9090"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	err := container.Invoke(func(app *iris.Application) {
		app.Run(
			iris.Addr(":"+port),
			iris.WithoutVersionChecker,
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
		)
	})

	if err != nil {
		panic(err)
	}
}
