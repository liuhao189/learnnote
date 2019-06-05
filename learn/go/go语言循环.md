# go语言循环语句

不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。

go语言的For循环有3种形式，只有其中的一种使用分号。

for init;condition;post{}

for condition{}

for{}

for循环的range格式可以对slice，map，数组，字符串等进行迭代循环。

```go
package main

import "fmt"

func main(){
	var b int=15;
	var a int;
	numbers:=[6]int{1,2,3,5};
	for a:=0;a<10;a++{
		fmt.Printf("a is %d\n",a);
	}
	fmt.Printf("************************");
	for a<b{
		a++;
		fmt.Printf("a is %d\n",a);
	}

	for i,x:=range numbers{
		fmt.Printf("第 %d 位x的值 = %d\n",i,x)
	}
}
```

## 循环控制语句

循环控制语句可以控制循环体内语句的执行过程，go支持break，continue，goto语句。

