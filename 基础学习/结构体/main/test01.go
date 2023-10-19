package main

import (
	"encoding/json"
	"fmt"
)

/*
定义Book结构体
*/
type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func test01() {
	fmt.Println("******创建一个新的结构体*************")
	fmt.Println(Book{"go", "aidan", "test", 12})
	// 也可以使用 key => value 格式
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com"})

	fmt.Println("********访问结构体成员*****************")
	var Book1 Book
	var Book2 Book

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.bookId)

	fmt.Println("*******结构体作为函数参数**************")
	printBook(Book1)
	printBook(Book2)

	fmt.Println("*******结构体序列化**************")

	p := Book{title: "go语言学习", author: "aidan", subject: "dd", bookId: 1}
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("错误信息", err.Error())
	}
	fmt.Println("序列化json数据=", string(jsonBytes))
}

func printBook(book Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
