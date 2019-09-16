package main

import "github.com/kataras/iris"

func commonHandler(ctx iris.Context) {
	request := ctx.Request()
	method := request.Method
	path := request.URL.Path
	ctx.Writef("Method is %s, path is %s.", method, path)
}

func main() {
	app := iris.Default()

	v1 := app.Party("v1")
	{
		v1.Post("/login", commonHandler)
		v1.Post("/submit", commonHandler)
		v1.Post("/read", commonHandler)
	}

	v2 := app.Party("/v2")
	{
		v2.Post("/login", commonHandler)
		v2.Post("/submit", commonHandler)
		v2.Post("/read", commonHandler)
	}

	app.Run(iris.Addr(":8080"))
}
