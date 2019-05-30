# Go语言变量

变量来源于数学，是计算机语言中能存储计算结果或能表示值抽象概念。

字母数字下划线，声明变量的一般形式是使用var关键字。

```go
var identifier1, identifier2 type
```

指定变量类型，如果没有初始化，则变量默认为零值。数值类型为0，布尔类型为false，字符串为""，引用类型为nil。

可以自动推断变量类型，根据赋值的字面量值。

省略var，:=左侧如果没有声明新的变量，就产生编译错误。

类型相同的多个变量，非全局变量。

```go
var vname1,vname2,vname2 type;
var vname1,vname2,vname3 = v1, v2, v3
```

```go
package main

import "fmt";

func main(){
	var a string;
	var b int;
	var c bool;
	var d *int;
	var e=true;
	f:=666;
	var h,i,j=9,8,7;
	fmt.Println(a,b,c,d,e,f);
	fmt.Println(h,i,j);
}

func main(){
	var a string="Liuhao"
	fmt.Println(a);
	var b,c int =1,2;
	fmt.Println(b,c);
}
```

## 值类型和引用类型

所有像int，float，bool和string这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值。

使用等号将一个变量赋值给另一个变量时，实际上是内存中将i的值进行了拷贝。

可以通过&i获取变量i的地址，值类型的变量值存储在栈中。

更复杂的数据通常使用多个字，这些数据一般使用引用类型保存。引用类型的变量存储的是r1的值所在的内存地址，或内存地址中第一个字所在的位置。

这个内存地址称值为指针，指针实际上也被存在另外的某一个字中。

赋值时，只有引用地址被复制。

## 简单形式

使用:=赋值操作符，使用变量的首选形式，它只能被用在函数体内，而不可以用于全局变量的声明与赋值。

局部变量不允许声明后不使用，全局变量允许声明但不使用。

空白标识符_实际上一个只写变量，不能得到它的值。Go语言中你必须使用所有被声明的变量，有时并不需要从一个函数得到的所有返回值。

```go
package main

import "fmt";

var k=999;

func main(){
	var i=666;
	// var i=555;
	k:=777;
	// j:=888;
	fmt.Println(&i,i,k);
}
```

# go语言常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量。常量中的数据类型只可以是布尔型，数字型和字符串型。
可以省略类型说明符[type]，因为编译器可以根据变量的值来推断其类型。

```go
const indentifier [type]= value
```

常量还可以用作枚举。常量就可以用len，unsafe.Sizeof函数计算表达式的值，常量表达式中，函数必须是内置函数，否则编译不过。

iota，特殊常量，可以认为是一个可以被编译器修改的常量。iota在const关键字出现时将被重置为0，每新增一行常量声明将使iota计数一次。

字符串类型在go里是个结构，包含指向底层数组的指针和长度，两个部分每部分都是8个字节。所以字符串类型大小为16个字节。

定义常量时，如果不提供初始值，则表示将使用上行的表达式。

```go
package main

import "fmt"
import "unsafe"

const (
	Unknown=0;
	Female=1;
	Male=2;
)

const (
	e="abc";
	f=len(e);
	g=unsafe.Sizeof(e);
)

func main(){
	const LENGTH int=10;
	const WIDTH int=5;
	var area int;
	const a,b,c=1,false,"str";
	area= LENGTH * WIDTH;
	fmt.Printf("Area is %d\n",area);
	fmt.Println(a,b,c);
	fmt.Println(Unknown,Female,Male);

	fmt.Println(e,f,g);
}

func main(){
	const (
		a=iota;
		b;
		c;
		d="ha";
		e;
		f=100;
		g;
		h=iota;
		i
	)

	fmt.Println(a,b,c,d,e,f,g,h,i)
}
```

