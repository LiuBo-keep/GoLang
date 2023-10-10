package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone *NokiaPhone) call() {
	fmt.Println("my is nokia Phone.....")
}

type IPhone struct {
}

func (iphone *IPhone) call() {

	fmt.Println("my is Iphone.....")
}

func test01() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
