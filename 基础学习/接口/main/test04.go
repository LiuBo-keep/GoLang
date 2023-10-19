package main

import "fmt"

type Controller struct {
	M int64
}

type Something interface {
	Get()
	Post()
}

func (c *Controller) Get() {
	fmt.Print("GET")
}

func (c *Controller) Post() {
	fmt.Print("POST")
}

type T struct {
	Controller
}

func (t *T) Get() {
	fmt.Print("T")
}

func (t *T) Post() {
	fmt.Print("T")
}

func test04() {

	var s Something

	s = new(T)

	var t T
	t.M = 1

	s.Get()
}
