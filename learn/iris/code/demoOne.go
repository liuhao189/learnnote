package main

import "github.com/kataras/iris"

type User struct {
	Username  string `Json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request fro path : %s", ctx.Path())
		ctx.Next()
	})

	app.Done(func(ctx iris.Context) {})

	app.Post("/decode", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s is %d years old and come from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
