package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("a 变量的地址:%x\n", &a)
	var ip *int
	ip = &a
	fmt.Printf("a 变量的地址是:%x\n", ip)
	fmt.Printf("a is %d\n", *ip)

	const MAX int = 3
	arr := []int{10, 100, 200}
	var ptr [MAX]*int
	var i, j int
	for i = 0; i < MAX; i++ {
		ptr[i] = &arr[i]
	}

	for j = 0; j < MAX; j++ {
		fmt.Printf("a[%d] = %d\n", j, *ptr[j])
	}

	pp()

	testSwap()
}

func testSwap() {
	var a int = 100
	var b int = 200
	fmt.Printf("A is %d\n", a)
	fmt.Printf("B is %d\n", b)
	swap(&a, &b)
	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func pp() {
	var a int
	var ptr *int
	var pptr **int
	a = 3000

	ptr = &a
	pptr = &ptr

	fmt.Printf("a is %d\n", a)
	fmt.Printf("*ptr is %d\n", *ptr)
	fmt.Printf("**pptr is %d\n", **pptr)
}
