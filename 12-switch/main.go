package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welocome to switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")

	case 2:
		fmt.Println("Dice value is one and you can open")

	case 3:
		fmt.Println("Dice value is one and you can open")

	case 4:
		fmt.Println("Dice value is one and you can open")

	case 5:
		fmt.Println("Dice value is one and you can open")

	case 6:
		fmt.Println("Dice value is one and you can open")

	default:
		fmt.Println("Nothing")
	}

}
