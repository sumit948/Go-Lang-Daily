package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	//declaring pointer
	// var mypointer *int
	// fmt.Println(mypointer)

	myNumber := 10

	var myptr = &myNumber

	fmt.Println("check pointer: ", myptr)
	fmt.Println("check pointer: ", *myptr)

	//we can manipulate pointer
	*myptr = *myptr + 20
	fmt.Println("new pointer: ", myNumber)
}
