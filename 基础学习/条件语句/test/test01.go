package test

import "fmt"

func Test01() {
	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)

	if second <= 8 {
		var gender string
		fmt.Println("请输入性别")
		fmt.Scanln(&gender)

		if gender == "男" {
			fmt.Println("进入决赛的男子组")
		} else {
			fmt.Println("进入决赛的女子组")
		}

	} else {
		fmt.Println("out....")
	}
}
