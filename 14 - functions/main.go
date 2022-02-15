package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	fmt.Println("Addition of a and b: ", addition(2, 2))
	fmt.Println("Pro addition", proaddition(2, 2, 2, 2, 2))
}
func addition(a int, b int) int {
	return a + b
}

func proaddition(values ...int) int {
	sum := 0

	for _, val := range values {
		sum += val

	}
	return sum
}
