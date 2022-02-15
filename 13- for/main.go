package main

import "fmt"

func main() {
	fmt.Println("Welcome to for loop")

	myarr := []string{"Sunday", "Monday", "Tuesday"}

	// for i := range myarr {
	// 	fmt.Println(myarr[i])

	// }

	for _, day := range myarr {
		//fmt.Printf("indext is %v and value is %v\n", index, day)
		fmt.Printf("indext is  and value is %v\n", day)
	}

}
