package main

import "fmt"

func main() {
	fmt.Println("----------------for 循环------------------")
	// 正常
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init 和 post 参数是可选的，我们可以直接省略它，类似 While 语句。
	sum1 := 1
	for sum1 < 10 {
		sum1 += 1
	}
	fmt.Println(sum1)

	// sum := 0
	//   for {
	//      sum++ // 无限循环下去
	//   }
	//   fmt.Println(sum) // 无法输出

	fmt.Println("----------------For-each range 循环------------------")
	// 这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。
	arry := []string{"google", "go"}

	// 使用普通方式
	for i := 0; i < len(arry); i++ {
		fmt.Print(arry[i])
	}

	// 使用foreach方式(i是下标，s是值，如果不想要下标可以使用_替代i)
	for i, s := range arry {
		fmt.Println(i, s)
	}

	maps := make(map[int]string)
	maps[1] = "李四"
	maps[2] = "王五"
	maps[3] = "赵六"
	// 读取 key 和 value
	for k, v := range maps {
		fmt.Println(k, v)
	}
	// 读取 key
	for key := range maps {
		fmt.Printf("key is: %d\n", key)
	}
	// 读取 value
	for _, value := range maps {
		fmt.Printf("value is: %f\n", value)
	}

	fmt.Println("----------------break------------------")

	for i := 0; i < 10; i++ {
		if i == 6 {
			fmt.Println("退出...")
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----------------continue------------------")
	for i := 0; i < 10; i++ {
		if i == 6 {
			fmt.Println("跳过6...")
			continue
		}
		fmt.Println(i)
	}

}
