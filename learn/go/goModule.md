# 开始使用Go Module

go的包管理方式是逐渐演进的，最初是monorepo模式，所有的包都放在GOPATH里面，使用类似命名空间的包路径区分包，不过这种包管理有问题，由于包依赖可能会引入破坏性更新，生产环境和测试环境会出现运行不一致的问题。

从V1.5开始引入vendor包模式，如果项目目录下有vendor目录，那么go工具链会优先使用vendor内的包进行编译，测试等。第三方的包管理思路都是通过这种方式来实现的，比如说社区维护准官方包管理工具dep。

不过官方并不认同这种方式，在V1.11中加入了Go Module作为官方包管理形式，dep无奈的结束了使命。

v1.11和v1.12的go版本中的gomod是不能直接使用的。通过go env命名返回值的GOMOD字段是否为空来判断是否已经开启了gomod，没有开启，通过设置环境变量开启。

gomod在Go v1.12功能基本稳定，到下一个版本v1.13将默认开启。

# Creating a Module

我们创建自己的包，testmod。

文件夹应该在$GOPATH之外，因为在GOPATH内默认不支持module的。

```bash
mkdir testmod
cd testmod
go mod init github.com/liuhao189/testmod
git init
git add *
git commit -am "first commit"
git push -u origin master
```
所有人都可以通过go get获取这个包。

这回拉取master分支的最新代码，但是可能会有破坏性更新。

## module版本

go modules有版本的。Go会使用tags作为版本。当我们发布包的时候，我们需要给我们的代码仓库打tag。

创建一个版本分支以便修改该版本的代码错误。

```bash
git tag v1.0.0
git push --tags
git checkout -b v1
git push -u origin v1
```

## 使用module

```go
package main

import "fmt"
import "github.com/liuhao189/gomodtest"

func main() {
	fmt.Println(gomodtest.Hi("liuhao"))
}

```

go build 命令会自动获取被程序引用的包。go.mod文件会加入程序引用的包。

go.sum包含了module的hash值们。

```bash
module main

go 1.12

require github.com/liuhao189/gomodtest v1.0.0
```

```bash
github.com/liuhao189/gomodtest v1.0.0 h1:F5k9CkylbSB+yWfOtI0UdgdjQvVubt+c5A1/We3LQLA=
github.com/liuhao189/gomodtest v1.0.0/go.mod h1:5UwVAv4ss7I12grCQ/DG10F6wcpt6paf/rPK7A7Iwa4=
```

## 创建一个bugfix的版本

```bash
# change your code base on version branch
git tag v1.0.1
git push --tags origin v1
 ```
使用go get -u packagename来更新package版本。

 ```bash
 go get -u 更新minor和patch更新
 go get -u=patch使用最新的patch更新。
 go get package@version 更新到特定的version
 ```
## 大版本更新

大版本可能不会兼容之前版本。一个大版本是一个完全不同的包。

我们的API不再和V1.X版本兼容，我们需要升级版本到V2。

V2和以上的版本会改变import的path。module github.com/liuhao189/testgomod/v2

```go
package gomodtest

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string,error) {
	swicth lang{
	case "en":
		return fmt.Sprintf("Hi,%s!",name),nil;
	case "pr":
		return fmt.Sprintf("Oi,%s!",name),nil;
	case "es":
		return fmt.Sprintf("!Hola,%s!",name),nil;
	case "fr":
		return fmt.Sprintf("Bonjour,%s!",name),nil;
	default:
		return "",errors.New("unknown language")
	}
}

```

```go
package main

import "fmt"
import gomodV2 "github.com/liuhao189/gomodtest/v2"
import "github.com/liuhao189/gomodtest"

func main() {
	fmt.Printf(gomodtest.Hi("liuhao") + "\n")
	g, err := gomodV2.Hi("liuhao", "zh")
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
		return
	}
	fmt.Println(g)
}

```
 
当我运行go build时，go会获取version 2.x.x的包。即使包名最后以v2结尾，依然需要使用gomodtest引用包。

我们可以同时引用一个包的不同版本。

## 整理包依赖

Go不会主动移除go.mod中的依赖。如果你需要移除不再使用的包，需要使用go mod tidy命令。

## vendor

go模块默认忽略vendor文件夹。如果我们也想把第三方包添加到我们的代码仓库中。我们可以使用go mod vendor。

创建一个vendor文件夹，并把所有的依赖拷贝到vendor文件夹中。

build到默认忽略vendor文件夹，需要使用go build -mod vendor来使用vendor文件夹。

```bash
go mod vendor
go build -mod vendor
```

## 结论

go module的使用起来很容易。开发者只需要import包，go会安装包并引用包。

当我们编译时，依赖会自动下载。同时，也宣告了$GOPATH的灭亡。