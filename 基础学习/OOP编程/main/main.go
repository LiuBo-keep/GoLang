package main

import "go_code/基础学习/OOP编程/employee"

func main() {
	e := employee.New("aidan", "liu", 12, 12)
	e.LeavesRemaining()
}
