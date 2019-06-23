package main

import "fmt"
import gomodV2 "github.com/liuhao189/gomodtest/v2"

func main() {
	g, err := gomodV2.Hi("liuhao", "zh")
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
		return
	}
	fmt.Println(g)
}
