# go语言入门

Go是一个开源的编程语言，它能让构造简单，可靠且高效的软件变得容易。Go有Linux，Mac和Windows等系统下的二进制发行版。

Go语言具有很强的表达能力，它简洁，清晰而高效。得益于其并发机制，用它编写的程序能够非常有效地利用多核与联网的计算机。其新颖的类型系统则使程序结构变得灵活而模块化。Go代码编译成机器码不仅非常迅速，还具有方便的垃圾收集机制和强大的运行时反射机制。

快速的，静态类型的编译型语言，感觉却像动态类型的解释型语言。

## go语言特色

简洁，快速，安全；并行，有趣，开源；内存管理，数组安全，编译迅速。

## go语言用途

go语言被设计成一门应用于搭载web服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。高性能分布式系统领域而言，Go语言比大多数其它语言有着更高的开发效率。提供海量并行的支持。

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello,World!")
}
```

```go
go run hello.go

go build hello.go
```
## go语言结构

包声明，引入包，函数，变量，语法&&表达式，注释。

源文件中非注释的第一行指明这个文件属于哪个包，package main表示一个可独立执行的程序，每个go应用程序都包含一个名为main的包。

main函数是每一个可执行程序所必须包含的，启动后第一个执行的函数，如果有init函数则会先执行该函数。

单行注释以//开头，多行注释以/*开始，*/结束。

当标识符，包括常量，变量，类型，函数名，结构字段以一个大写字母开头，这种形式的标识符的对象就可以被外部包的代码所使用(public)，这被称为导出。标识符如果以小写字母开头，则对包外是不可见的，在整个包的内部是可见并且可用的(protected)。

{不能放在单独的一行。

```go
package main

import "fmt"

/* then run main func */
func main(){
	fmt.Println("Hello,World!")
}

/* run init func first */
func init(){
	fmt.Println("Init!")
}
```

# go语言基础语法

## 行分隔符

Go程序中，一行代表一个语句结束，每个语句不需要以分号;结尾。这些工作都将由Go编译器自动完成。多个语句写在同一行，必须以分号区分。

```go
package main

import "fmt"

func init(){
	fmt.Println("Init func");fmt.Println("The second line!")
}

func main(){
	fmt.Println("Main func!");
}
```

## 注释

注释不会被编译，每一个包应该有相关注释。单行注释是最常见的注释形式，可以在任何地方以//开头的单行注释。多行注释也叫块注释，均以/*开头，以*/结尾。

## 标识符

标识符用来命名变量，类型等程序实体。一个标识符实际上就是一个或多个字母，下划线组成的序列，第一个字符必须是字母或下划线。

## 关键字

break；default；func；interface；select；case；defer；go；map；struct；chan；else；goto；package；switch；const；fallthrough；if；range；type；continue；for；import；return；var；

预定义标识符：append，bool，byte；cap，close，complex，complex64，complex128，unit16，copy，false，float32，float64，imag，int，int8，int16，unit32，int32，int64，iota，len，make，new，nil，panic，uint64，
print，println，rea，recover，string，true，uint，uint8，uintptr。

## 语言的空格

Go语言中变量的声明必须使用空格隔开。语句中适当使用空格能让程序更易阅读。

```go
var age int;
```

# go语言数据类型

go编程语言中，数据类型用于声明函数和变量。数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，可以充分利用内存。

## 布尔型

布尔型的值只可以是常量true或false，var b bool = true

## 数字类型

整型int和浮点数float32，float64，Go语言支持整型和浮点型数字，并且支持复数，其中为的运算采用补码。

| uint8 | 0-255  | int8 | -128-127 |

| uint16 | 0-65536 | int16 | -32768-32767 |

| uint32 | 0- 42亿 |  int32 | -21亿-21亿 | 

| uint64 | 0- 1亿万亿 | int64 | 很大 |

| byte 类似于uint8 |
| rune 类似于int32 |
| uint 32或64位 |
| int  与uint一样大小 |
| uintptr 无符号整型，用于存放一个指针 |

## 字符串类型

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本。

## 派生类型

指针类型Pointer，数组类型，结构化类型struct，Channel类型，函数类型，切片类型，接口类型，Map类型。

