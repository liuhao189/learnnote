# Iris

iris框架为你的下一个网站，API，或分布式应用程序提供一个简单且有效，并且易于使用的基础。

# 安装

唯一的环境要求是安装了GO语言环境，最低版本要求是1.8版本，但是强烈推荐使用1.102。

# Quick Start

```go
package code

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	//New no middleware
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
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

```

# Using Get Post Put Patch Delete and Options

支持以MethodName开头的方法，直接跟url pattern即可，后面是运行函数。

## Parmeters in path

1、:string  Params.Get

2、:int     depends on the host arch Params.GetInt

3、:intX[8,16,32,64]    Params.GetIntX

4、:uintX[8,16,32,64]    Params.GetUintX

5、:bool 1,"t","TRUE","true" 0,"f","F","FALSE","false","False" Params.GetBool

6、:alphabetical lowercase or uppercase letters  Params.Get

7、:file  lowercase or uppercase letters,numbers,underscore,dash,point and no spaces or other special characters that are not valid for filenames Params.Get

8、:path  can be sperated by slashes but should be the last part of the route path Params.Get

## Built-in Func

1、regexp  expx string

2、prefix prefix string

3、suffix  suffix string

4、contains s string

5、min minValue

6、max maxValue

7、range minValue maxValue

## do it yourselft

RegisterFunc可以接受任何返回func(paramValue string) bool的函数。或者func(steing) bool。

如果验证失败，它会返回404或者任意status code。

```go
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
```


