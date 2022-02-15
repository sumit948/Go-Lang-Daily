package main

import "fmt"

func main() {
	var userName string = "Jhon Weak"
	fmt.Println("Welcome: ", userName)
	fmt.Printf("Type of this username is: %T \n", userName)

	var isavailable bool = true
	fmt.Println("is avaliable: ", isavailable)
	fmt.Printf("type of isavilable: %T \n", isavailable)

	var smallFloat float32 = 150.4423561545544554454554
	var bigFloat float64 = 150.4423561545544554454554
	fmt.Println("float32: ", smallFloat)
	fmt.Println("float64: ", bigFloat)
	fmt.Println("__________________________________________________________")
	//default values and some alises
	var cheakString bool
	fmt.Println(cheakString)
	fmt.Println("_________________________________________________________")
	//implecit type
	var mystring = "this is my string"
	fmt.Println(mystring)

	//no var keyword

	number1 := 10
	fmt.Println("without var number1: ", number1)
	string1 := "Second string without var"
	fmt.Println(string1)

}
