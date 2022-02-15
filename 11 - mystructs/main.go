package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	//No Inheritance in go lang, super or parent

	check := User{"Sumit", "Kolhapur", "Sumit12@gmail.com", true}
	fmt.Println(check)
	fmt.Printf("Sumit Details: %+v\n", check)
}

type User struct {
	Name   string
	Add    string
	Email  string
	Status bool
}
