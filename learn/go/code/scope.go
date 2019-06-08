package main

import "fmt"

var g int

func main() {
	var a, b, c int
	a = 10
	b = 20
	c = a + b
	g = a + b
	fmt.Printf("Result is a = %d, b= %d and c=%d\n", a, b, c)
	fmt.Printf("g is %d\n", g)
}
