package main

import "github.com/kataras/iris"

import "github.com/kataras/iris/hero"

import "fmt"

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func login(form LoginForm) string {
	return "Hello  " + form.Username
}

func main() {
	app := iris.Default()
	//
	hero.Register(func(ctx iris.Context) (form LoginForm) {
		//bind the form with the x-www-form-urlencoded form data
		ctx.ReadForm(&form)
		fmt.Printf("username is %s, password is %s!\n", form.Username, form.Password)
		return
	})
	//
	loginHandler := hero.Handler(login)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<form action='login' method='post'><input name='username'/><input type='password' name='password'/><button type='submit'>submit</button></form>")
	})

	app.Post("/login", loginHandler)
	//
	app.Run(iris.Addr(":8080"))
}
