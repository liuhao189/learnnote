pakcage main

import "github.com/kataras/iris"


func main(){
	app:=iris.Default();

	app.Get('/someGet',func(ctx iris.Context){
		ctx.WriteString("some get!")
	})

	app.Post('/somePost',func(ctx iris.Context){
		ctx.WriteString("some post!")
	})

	app.Put('/somePut',func(ctx iris.Context){
		ctx.WriteString("some put!")
	})

	app.Delete("/someDelete",func(ctx iris.Context){
		ctx.WriteString("some Delete!")
	})

	app.Patch("/somePatch",func(ctx iris.Context){
		ctx.WriteString("some Patch!")
	})

	app.Head("/someHead",func(ctx iris.Context){
		ctx.WriteString("some Head!")
	})

	app.Options("/someOptions",func(ctx irir.Context){
		ctx.WriteString("some Options!")
	})

	//
	app.Get("/users/{id:uint64}",func(ctx iris.Context){
		id:=ctx.Params().GetUint64Default("id",0);
		ctx.WriteString("UserId is "+id)
	})
	//
	app.Get("/profile/{name:alphabetical max(255)}",func(ctx iris.Context){
		name:=ctx.Params().get("name");
		ctx.WriteString(name);
	})
	//
	app.Run(iris.Addr(":8080"))
}