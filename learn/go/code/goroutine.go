package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	fmt.Println("invoked!")
	for i := range ch {
		fmt.Println(i)
	}
}

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}
// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c
// 	fmt.Println(x, y, x+y)
// }
