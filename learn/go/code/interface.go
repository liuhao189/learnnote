package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
	price float32
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	name string
}

func (iphpne IPhone) call() {
	fmt.Println("I am Iphone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()
}
