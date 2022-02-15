package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Welcome := "Welcome to the Web"
	fmt.Println(Welcome)
	fmt.Println("Give rating 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you: ", input) //4
	numberrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Thank You: ", numberrating+1) //4+1 = 5
	}

}
