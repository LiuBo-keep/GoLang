package main

import "fmt"

// 接口断言
type I interface {
	print()
}

type P1 struct {
}

func (p1 *P1) print() {
	fmt.Println("p1")
}

type P2 struct {
}

func (p2 *P2) print() {
	fmt.Println("p2")
}

func sw(i interface{}) {

	//switch ins := i.(type) {
	//case P1:
	//	ins.print()
	//case P2:
	//	ins.print()
	//default:
	//	fmt.Println("not find")
	//}

	s, f := i.(P1)
	fmt.Println(f, s)
}

func test05() {

	sw(new(P1))
}
