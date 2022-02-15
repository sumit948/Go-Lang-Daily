package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("HEllo")
	greeter("world")
}

func greeter(name string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(name)
	}
}
