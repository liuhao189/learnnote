# go语言函数

函数是基本的代码块，用于执行一个任务。go语言最少有个main函数，可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明告诉编译器函数的名称，返回类型和参数。

go语言标准库提高了多种可用的内置函数。

## 函数定义

parameter list参数列表，参数就行一个占位符，当函数被调用时，可以将值传递给参数。参数列表指定的是参数类型、顺序、及参数个数。

return_types，返回类型。有些功能不需要返回值，这种情况下，return_types不是必须的。

函数可以返回多个值。


```go
func function_name([parameter list]) [return_types] {

}
```

## 函数参数

函数如果使用参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数，值传递，引用传递。默认情况下，Go语言使用的是值传递。

## 函数用法

go语言可以很灵活的创建函数，并作为一个变量。

go语言支持匿名函数，可作为必报，匿名函数是一个内联语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

go语言同时有函数和方法，一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者一个指针，所有给定类型的方法属于该类型的方法集。

```go
func (variable_name variable_data_type) function_name() [return_type]{

}
```

```go
package main 

import "fmt";
import "math";

func max(num1,num2 int) int{
	if(num1>num2){
		return num1;
	}else{
		return num2;
	}
}

func swap(x,y string) (string,string){
	return y,x;
}

func getSequence() func() int{
	i:=0;
	return func() int{
		i++;
		return i;
	}
}

type Circle struct{
	radius float64;
}

func (c Circle) getArea() float64{
	return 3.14 * c.radius * c.radius;
}

func main(){
	var maxNum=max(10,15);
	fmt.Printf("Max Num is %d\n",maxNum);
	a,b:=swap("Google","Bing");
	fmt.Println(a);
	fmt.Println(b);
	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9));

	nextNumber:=getSequence();
	fmt.Println(nextNumber());
	fmt.Println(nextNumber());
	fmt.Println(nextNumber());

	nextNumber1:=getSequence();
	fmt.Println(nextNumber1());
	fmt.Println(nextNumber1());

	fmt.Println("****************")

	var c1 Circle;
	c1.radius=10.00;
	fmt.Printf("Area is %f.\n",c1.getArea());
}
```