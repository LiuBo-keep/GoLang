package main

import "fmt"

type User struct {
}

func (u User) Error() string {
	return "not find user"
}

func main() {

	user := User{}

	fmt.Println(user.Error())
}
