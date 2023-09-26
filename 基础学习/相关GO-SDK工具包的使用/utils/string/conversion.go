package string

import (
	"fmt"
	"strconv"
)

// 统计字符串长度
func Len(str string) int {
	return len(str)
}

// 字符串转换整数
func StrToNumber(str string) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("转换错误：", err)
	}
	return n
}

// 整数转字符串
func NumberToStr(num int) string {
	return strconv.Itoa(num)
}

// 字符串转byte数组
func StrToByteArr(str string) []byte {
	return []byte(str)
}
