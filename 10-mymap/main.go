package main

import "fmt"

func main() {
	fmt.Println("Welcome to map")

	languages := make(map[int]string)

	languages[1] = "JavaScript"
	languages[2] = "Go lang"
	languages[3] = "Java"
	languages[4] = "Python"

	fmt.Println("map: ", languages)

	fmt.Println("Java key: ", languages[3])

	//-----for delete the key
	delete(languages, 4)
	fmt.Println("map: ", languages)
}
