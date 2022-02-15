package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	checkDefer()
}

//if defer in outside function
func checkDefer() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
