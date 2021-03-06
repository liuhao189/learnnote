# go语言运算符

运算符用于在程序运行时执行数学或逻辑运算。算术运算符，关系运算符，逻辑运算符，位运算符，赋值运算符，其它运算符。

算术运算符，+、-、*、/、%、++、--。

关系运算符，==、!=、>,<,>=,<=。

逻辑运算符，&&，||，！。

位运算符，&，|，^,<<,>>。

赋值运算符，=，+=，-=，*=，/=，%=，<<=,>>=,&=,^=,|=。

其它运算符，&返回变量存储地址，*指针变量。

二元运算符均是从左到右。

```go
package main

import "fmt";

func print(c int){
	fmt.Printf("val is %d\n",c);
}

func main() {
	var a int =21;
	var b int =10;
	var c int;
	c=a+b;
	print(c);
	c=a-b;
	print(c);
	c=a*b;
	print(c);
	c=a/b;
	print(c);
	c=a%b;
	print(c);
	a++;
	print(a);
	a=21;
	a--;
	print(a);
	var a int=21;
	var b int=10;
	if(a==b){
		fmt.Printf("a == b is true\n")
	}else{
		fmt.Printf("a == b is false\n")
	}
	if(a<b){
		fmt.Printf("a < b is true \n")
	}else{
		fmt.Printf("a < b is false \n")
	}

	if(a>b){
		fmt.Printf("a > b is true\n")
	}else{
		fmt.Printf("a>b is false\n")
	}

	var a bool=true;
	var b bool=false;

	if(a&&b){
		fmt.Printf("a && b is true\n");
	}else{
		fmt.Printf("a && b is false!\n");
	}

	if(a || b){
		fmt.Printf("a || b is true!\n");
	}

	a=false;
	b=true;
	if(!(a&&b)){
		fmt.Printf("!(a&&b) is true!");
	}

	var a uint = 60;//0011 1100 
	var b uint = 13;//0000 1101
	var c uint = 0;// 0000 0000

	c= a &b;
	fmt.Printf("c is %d\n",c);//12
	c= a| b;
	fmt.Printf("c is %d\n",c);//61
	c=a ^b;
	fmt.Printf("c is %d\n",c);//49
	c=a<<2;
	fmt.Printf("c is %d\n",c)//240
	c=a>>2;
	fmt.Printf("c is %d\n",c);//15

	var a =4;
	var b int32;
	var c float32;
	var ptr *int;
	fmt.Printf("a type is %T\n",a);
	fmt.Printf("b type is %T\n",b);
	fmt.Printf("c type is %T\n",c);
	ptr=&a;
	fmt.Printf("a value is %d\n",a);
	fmt.Printf("*ptr is %d\n",*ptr);
}
```

# go语言条件语句

条件语句需要开发者通过制定一个或多个条件，并通过测试条件是否为true来决定是否执行指定语句，并在条件为false的情况下执行另外的语句。

if语句，if...esle语句，if嵌套语句跟JS类似。

switch语句，基于不同条件执行不同动作。

每一条case都是唯一的，从上至下逐一测试，直到匹配为止。匹配项后面也不需要再加break。需要执行后面的case，可以使用fallthrough。

swicth后表达式可以是任何类型，case值是同类型的任意值，不局限于常量或整数。同时测试多个可能符合条件的值，逗号分割它们即可。

```go
package main

import "fmt";

func main() {
	var grade string ="B";
	var marks int = 90;
	switch marks{
	case 90:grade="A";
	case 80:grade="B";
	case 50,60,70:grade="C";
	default:grade="D";
	}

	switch{
	case grade == "A":
		fmt.Printf("优秀\n");
	case grade=="B":
		fmt.Printf("良好\n");
	case grade=="C":
		fmt.Printf("及格\n");
	case grade=="D":
		fmt.Printf("不及格\n");
	}
	fmt.Printf("Your grade is %s\n",grade);
}
```

## Type Switch

switch语句还可以被用于type-switch来判断某个interface变量中实际存储的变量类型。

## fallthrough

使用fallthrough会强制执行后面的case语句，fallthrough不会判断下一条的表达式结果是否为true。

```go
func main(){
	var x interface{}
	switch i:=x.(type){
	case nil:
		fmt.Printf("x is :%T",i);
	case int:
		fmt.Printf("x is int!");
	case float64:
		fmt.Printf("float64");
	case func(int) float64:
		fmt.Printf("x is func(int):float64");
	case bool,string:
		fmt.Printf("x is bool,string");
	default:
		fmt.Printf("unknown type!")
	}

	switch{
	case false:
		fmt.Printf("1，false");
		fallthrough;
	case true:
		fmt.Printf("2,true");
		fallthrough;
	case false:
		fmt.Printf("3,false");
		fallthrough;
	case true:
		fmt.Printf("4,true");
	case false:
		fmt.Printf("5,false");
		fallthrough;
	default:
		fmt.Printf("6,default");
	}
}
```

## select语句

select是Go中的一个控制结构，类似于用于通信的switch语句。select随机执行一个可运行的case，没有case可运行，它将阻塞，直到有case可运行。一个默认的子句总是可运行的。

每个case都必须是一个通信。

所有channel表达式都会被求值。

所有被发送的表达式都会被求值。

如果任意某个通信可以执行，它就执行，其它被忽略。

多个case都可以运行，select会随机公平地选出一个执行，其它不会执行。

如果没有case可以运行，有defaul子句，则执行该语句，否则，阻塞，直到某个通信可以运行。
