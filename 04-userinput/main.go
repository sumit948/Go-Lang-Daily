package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter Your name: ")

	// // comma ok  //err err
	// //input(try) and _(catch)
	// input, _ := reader.ReadString('#') //read string till # then stop inputing
	// fmt.Println("Thank You: ", input)
	// fmt.Printf("Type Check: %T", input)
	defer sampleuserinput()
}

func sampleuserinput() {
	// var nums int64
	// fmt.Println("Enter Two numbers: ")
	// fmt.Scanln(&nums, &nums)
	// fmt.Println("Your number is : ", nums)

	num4 := []int64{2, 3, 4, 5, 6}
	var sum int64
	for i := 0; i < len(num4); i++ {
		sum += num4[i]
	}
	fmt.Println(sum)
}
