package main

import "fmt"

type myint int
type mystr string

func test01() {
	var i1 myint
	var i2 = 100
	i1 = 100

	fmt.Println(i1)
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name)
	fmt.Println(s1)
}
