package main

import "fmt"

type person struct {
	age  int
	name string
	a    bool
}

func test02() {
	// new方式,未设置初始值的，会赋予类型的默认初始值
	p := new(person)
	p.name = "李四"
	fmt.Println(p.age)  // 0
	fmt.Println(p.a)    // false
	fmt.Println(p.name) // 李四
}
