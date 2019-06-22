package main

import "fmt"
import "unsafe"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	e = "abc"
	f = len(e)
	g = unsafe.Sizeof(e)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("Area is %d\n", area)
	fmt.Println(a, b, c)
	fmt.Println(Unknown, Female, Male)

	fmt.Println(e, f, g)
}

func main() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
