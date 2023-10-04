package main

import (
	"errors"
	"fmt"
)

type User struct {
}

func (u User) Error() string {
	return "not find user"
}

func findUser(name string) (string, error) {

	if name != "李四" {
		return name, errors.New("用户信息错误.")
	}
	fmt.Println("ssssss")
	return name, nil
}

func main() {

	user := User{}

	fmt.Println(user.Error())

	s, err := findUser("1111")
	fmt.Println(err)

	fmt.Println(s)

}
