package main

import (
	"fmt"
	"time"
)

func main() {

	go test01()
	test01()

	fmt.Println("--------------------")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func test01() {

	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println(i)
	}
}

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}
