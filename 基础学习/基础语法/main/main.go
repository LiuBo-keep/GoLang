package main

import "fmt"

func main() {

	// 行分隔符
	fmt.Println("Hello world")
	fmt.Println("test")

	// 字符串拼接
	fmt.Println("Google" + "Ro")

	// 字符串格式化
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

}
