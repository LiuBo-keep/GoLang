package utils

import "fmt"

/*
将计算的功能，放到一个函数中。然后在需要使用，调用即可
为了让其他包的文件使用Cal函数，需要将c大小类型与其他语言的public
*/
func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64

	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符错误...")
	}

	return res
}
