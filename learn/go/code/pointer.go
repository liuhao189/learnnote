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

}
