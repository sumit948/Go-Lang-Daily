package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods")
	check := User{"Sumit", 21, true}
	fmt.Println(check)
	check.GetStatus()
}

type User struct {
	Name   string
	Age    int
	status bool
}

func (u User) GetStatus() {
	fmt.Println("Getting status: ", u.status)
}
