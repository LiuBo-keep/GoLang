package main

import (
	"fmt"
	string2 "go_code/基础学习/相关GO-SDK工具包的使用/utils/string"
	"go_code/基础学习/相关GO-SDK工具包的使用/utils/time"
)

var str string = "lisi"

func main() {

	fmt.Println("----------------字符串-----------------------")
	fmt.Println(string2.Len(str))
	fmt.Println(string2.StrToNumber(str))
	fmt.Println("----------------时间的-----------------------")
	fmt.Println(time.GetYear())
	fmt.Println(time.GetMonth())
	fmt.Println(time.GetDay())
	fmt.Println(time.GetHour())
	fmt.Println(time.GetMinute())
	fmt.Println(time.GetSecond())

}
