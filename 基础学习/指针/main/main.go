package main

import "fmt"

func main() {
	fmt.Println("*************简单使用**********")
	// 声明实际变量
	var c int = 10
	// 声明指针变量
	var prt *int

	prt = &c

	fmt.Println("a 变量的地址是: %x\n", &c)
	// 指针变量的存储地址
	fmt.Printf("ip 变量储存的指针地址: %x\n", prt)
	// 使用指针访问值
	fmt.Printf("*ip 变量的值: %d\n", *prt)

	// 通过指针改变a的值
	*prt = 20
	fmt.Println(c)

	fmt.Println("*************空指针**********")

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)

	fmt.Println("*************指针作为函数参数**********")
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
