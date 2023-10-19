package main

import "fmt"

type Human03 interface {
	Len()
}

type Student03 interface {
	Human03
}

type Test struct {
}

func (h *Test) Len() {
	fmt.Println("len.....")
}

func test03() {
	var s Student03
	s = new(Test)
	s.Len()
}
