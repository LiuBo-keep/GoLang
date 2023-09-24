package main

import "fmt"

func main() {
	fmt.Println("--------------函数返回两个数的最大值----------------------")
	fmt.Println(max(10, 15))

	fmt.Println("--------------函数返回多个值----------------------")
	name1, name2, name3 := names()
	fmt.Println(name1, name2, name3)

	fmt.Println("--------------函数参数(值传递)----------------------")
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)
	/* 通过调用函数来交换值 */
	swap1(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)

	fmt.Println("--------------函数参数(引用传递)----------------------")
	/* 定义局部变量 */
	var c int = 100
	var d int = 200

	fmt.Printf("交换前，a 的值 : %d\n", c)
	fmt.Printf("交换前，b 的值 : %d\n", d)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap2(&c, &d)

	fmt.Printf("交换后，a 的值 : %d\n", c)
	fmt.Printf("交换后，b 的值 : %d\n", d)
}

/*
函数返回两个数的最大值
*/
func max(number1, number2 int) int {
	var max int = 0

	if number1 > number2 {
		max = number1
	}
	if number1 < number2 {
		max = number2
	}

	return max
}

/*
函数返回多个值
*/
func names() (string, string, string) {
	return "lisi", "wangwu", "zhaoliu"
}

/*
定义相互交换值的函数(值传递)
*/
func swap1(x, y int) {
	var temp int
	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/
}

/*
定义交换值函数(引用传递)
*/
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
