package main

import "fmt"

type personAa struct {
	width, height int
}

func (p *personAa) setVal() {
	p.height = 20
}

func test03() {

	p := personAa{1, 2}
	s := p
	s.setVal()

	fmt.Println(p.height, s.height)
}
