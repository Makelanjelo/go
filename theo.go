package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
}

func main() {
	fmt.Println("Start programm!")

	// var user User = User{name: "john", age: 23, password: "pass"}
	user := User{"John", 23, "pass"}
	fmt.Println(user)
}
