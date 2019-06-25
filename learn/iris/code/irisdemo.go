package code

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().setLevel("debug")
	//
	app.Use(recover.New())
	app.Use(logger.New())
	//
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	//
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	//
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	//
	app.Macros().Get("string").RegisterFunc("range", func(minLength, maxLength int) func(string) bool {
		return func(paramValue string) bool {
			return len(paramValue) >= minLength && len(paramValue) <= maxLength
		}
	})
	//
	app.Get("/limitchar/{name: string range(1,200) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length 
		 otherwise this hanlder will not be executed`, name)
	})
	//
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
