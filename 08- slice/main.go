package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welocome to slice")

	var slice1 = []string{"apple", "samsung", "nokia"}
	fmt.Println("slice1: ", slice1)

	slice1 = append(slice1, "blackbeary")
	fmt.Println("New Slice1: ", slice1)

	// slice1 = append(slice1[1:3]) //start:end values
	// fmt.Println("After removing place 1: ", slice1)

	scores := make([]int, 4)
	scores[0] = 55
	scores[1] = 59
	scores[2] = 45
	scores[3] = 35
	//scores[0] = 85 can't add
	scores = append(scores, 89, 25, 7)
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println("sorted list: ", scores)

	//------------------------------------------------------------------
	//how to remove particular value in slice

	course := []string{"CSE", "Mech", "Chem", "Civil"}
	fmt.Println("course list: ", course)

	var indexvalue int = 2
	course = append(course[:indexvalue], course[indexvalue+1:]...)
	fmt.Println("new course list", course)

}
