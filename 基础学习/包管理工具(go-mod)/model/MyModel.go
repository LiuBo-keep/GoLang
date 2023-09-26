package model

import "fmt"

type Model struct {
	name string
}

func (m *Model) MyMode() {

	fmt.Println("我是MyModel方法")
}
