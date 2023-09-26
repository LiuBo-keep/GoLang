package main

import (
	"fmt"
	"go_code/基础学习/包/utils"
)

func main() {

	var n1 float64 = 1.2
	var n2 float64 = 6.4
	var operator byte = '+'

	res := utils.Cal(n1, n2, operator)

	fmt.Println("res=", res)
}
