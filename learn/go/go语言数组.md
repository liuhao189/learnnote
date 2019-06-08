# Go语言数组

go语言提供数组类型的数据结构。

## 声明数组

go语言数组声明需要制定元素类型及元素个数。

```go
var variable_name [SIZE] variable_type
```

## 初始化数组

初始化数组中{}中的元素个数不能大于[]中的数字。
可以忽略[]中的数组大小，编译器会自动推断。
```go
var balance = [5]float32{1000.0,2.0,3.0,3.4,7.0,50.0}
```
## 访问数组元素

数组元素可以通过索引来读取，格式为数组后加中括号，中括号中为索引的值。

## go多维数组

```go
var variable_name [SIZE1][SIZE2]....[SIZEN] variable_type
```


