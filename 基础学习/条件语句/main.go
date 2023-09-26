package main

import (
	"fmt"
	"go_code/基础学习/条件语句/test"
	"time"
)

var number int = 9

func main() {

	// if语句
	if number > 10 {
		fmt.Println("number 大于10")
	}

	fmt.Println("number 小于10")

	fmt.Println("------------if ... else 语句-------------")

	// if ... else 语句
	if number > 10 {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

	fmt.Println("---------- if嵌套---------------")

	// if嵌套
	if number < 10 {
		if number > 5 {

		}
	}

	fmt.Println("-----------switch--------------")

	// switch
	age := 12
	switch age {
	case 10:
		fmt.Println("age is 10")
	case 12:
		fmt.Println("age is 12")
	case 13:
		fmt.Println("age is 13")
	default:
		fmt.Println("age not")
	}

	fmt.Println("------------fallthrough-------------")
	// fallthrough
	switch {
	case false:
		fmt.Println("1")
		fallthrough
	case true:
		fmt.Println("2")
		fallthrough
	case false:
		fmt.Println("3")
		fallthrough
	case true:
		fmt.Println("4")
	case false:
		fmt.Println("5")
	default:
		fmt.Println("6")
	}

	fmt.Println("------------select-------------")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("22222")
		c1 <- "one"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("333333")
		c2 <- "tow"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("111111")
		}
	}

	fmt.Println("------------------test01------------------")
	test.Test01()
}
