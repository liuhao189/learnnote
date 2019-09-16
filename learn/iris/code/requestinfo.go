package main

import "github.com/kataras/iris"
import "mime/multipart"
import "strings"

const maxSize = 5 << 20

func main() {
	app := iris.Default()

	app.Get("/welcome", func(ctx iris.Context) {
		firsname := ctx.URLParamDefault("firstname", "Guest")
		lastname := ctx.URLParam("lastname")
		ctx.Writef("Hello %s %s", firsname, lastname)
	})

	app.Post("form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")
		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	app.Get("/form", func(ctx iris.Context) {
		ctx.HTML("<form action='upload' method='post' enctype='multipart/form-data'><input type='file' name='file'/><button type='submit'>submit</button></form>")
	})

	app.Post("/upload", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		ctx.UploadFormFiles("./uploads", beforeSave)
	})

	app.Run(iris.Addr(":8080"))
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)
	file.Filename = ip + "-" + file.Filename
}
