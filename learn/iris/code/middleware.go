package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/middleware/recover"
import "github.com/kataras/iris/middleware/logger"

func main() {
	app := iris.New()
	app.Use(recover.New())

	requestLogger := logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	})

	app.Use(requestLogger)
	//
	v1 := app.Party("v1")
	{
		v1.Get("/login", func(ctx iris.Context) {
			ctx.HTML("<h1>Login Page</h1>")
		})
	}
	//
	app.Run(iris.Addr(":8080"))

}
