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

每一条case都是唯一的，从上至下逐一测试，直到匹配为止。匹配项后面也不需要再加break。需要执行后门的case，可以使用fallthrough。

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