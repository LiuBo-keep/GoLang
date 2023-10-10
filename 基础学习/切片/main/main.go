package main

import "fmt"

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	// 修改切片
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// len() 和 cap() 函数
	var numbers = make([]int, 3, 5)
	numbers[2] = 1
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)

	// append() 和 copy() 函数
	var number []int
	printSlice(number)
	number = append(number, 0)
	printSlice(number)
	/* 向切片添加一个元素 */
	number = append(number, 1)
	printSlice(number)
	/* 同时添加多个元素 */
	number = append(number, 2, 3, 4)
	printSlice(number)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(number), (cap(number))*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, number)
	printSlice(numbers1)
}

func printSlice(number []int) {
	fmt.Println(number)
}
