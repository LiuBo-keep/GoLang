package main

import "fmt"

type HumanTest struct {
	name  string
	age   int
	phone string
}

type StudentTest struct {
	HumanTest
	school string
}

type EmployeeTest struct {
	HumanTest
	company string
}

func (h *HumanTest) SayHi() {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *EmployeeTest) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func test05() {
	mark := StudentTest{HumanTest{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := EmployeeTest{HumanTest{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
