## type关键字

type是go语法里的重要而且常用的关键字，type绝不只是对应于C/C++中的typedef。搞清楚type的使用，就容易理解go语言中的核心概念struct、interface、函数等的使用。

### 一、类型定义

#### 1.1 定义结构体

使用type 可以定义结构体类型：

```go
//1、定义结构体
//结构体定义
type person struct {
name string //注意后面不能有逗号
age  int
}
```

#### 1.2 定义接口

使用type 可以定义接口类型：

```go
type USB interface {
start()
end()
}

```

#### 1.3 定义其他的新类型

使用type，还可以定义新类型。

语法：
```go
type 类型名 Type

```

示例代码：
```go
package main

import "fmt"

type myint int
type mystr string

func main() {

	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1)
	//i1 = i2 //cannot use i2 (type int) as type myint in assignment
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name)
	fmt.Println(s1)
	name = s1 //cannot use s1 (type string) as type mystr in assignment
}

```

### 1.4 定义函数的类型

